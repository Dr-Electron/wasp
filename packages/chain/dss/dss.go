// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// Run a NonceDKG and sign the supplied hash.
//
// This is a simplified implementation.
// Later the DKG part can be run in advance, while waiting for transactions.
//
// The general workflow is the following:
//
//  1. Start it upon activation of a step (last stateOutput is approved).
//  2. Exchange the underlying messages until:
//     2.1. ACSS Intermediate output is received.
//  3. Then wait for the ACS and then the VM to complete:
//     3.1. pass the ACS result to the nonce-dkg (to complete the nonces).
//     3.2. pass the VM output as a message to sign (its hash).
//  4. Exchange messages until the signature is produced.
//  5. Output the signature.
//
// TODO: Make sure no two signatures are ever produced by the nonce-dkg for the same
//
//	base TX. That would reveal the permanent private key of the committee.
package dss

import (
	"fmt"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/sign/dss"
	"go.dedis.ch/kyber/v3/suites"

	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/gpa"
	"github.com/iotaledger/wasp/packages/gpa/adkg/nonce"
	"github.com/iotaledger/wasp/packages/tcrypto"
)

type DSS interface {
	AsGPA() gpa.GPA
}

type Output struct {
	ProposedIndexes []int  // Intermediate output.
	Signature       []byte // Final output.
}

const (
	subsystemDKG byte = iota
)

type dssImpl struct {
	suite                    suites.Suite
	withWrappers             gpa.GPA // This instance, with all the wrappers.
	me                       gpa.NodeID
	mySK                     kyber.Scalar
	nodeIDs                  []gpa.NodeID
	nodePKs                  map[gpa.NodeID]kyber.Point
	f                        int
	longTermSecretShare      tcrypto.SecretShare
	dkg                      gpa.GPA
	dkgOutIndexes            []int                // Intermediate DKG output.
	dkgDecidedIndexProposals map[gpa.NodeID][]int // ACS decision.
	dkgOutNonce              dss.DistKeyShare     // Final DKG output.
	messageToSign            []byte
	dssPartialSigBuffer      map[gpa.NodeID]*dss.PartialSig // Accumulate early partial signatures
	dssSigner                *dss.DSS
	signature                []byte // The output.
	msgWrapper               *gpa.MsgWrapper
	log                      *logger.Logger
}

var _ DSS = &dssImpl{}

func New(
	suite suites.Suite,
	nodeIDs []gpa.NodeID,
	nodePKs map[gpa.NodeID]kyber.Point,
	f int,
	me gpa.NodeID,
	mySK kyber.Scalar,
	longTermSecretShare tcrypto.SecretShare,
	log *logger.Logger,
) DSS {
	d := &dssImpl{
		suite:                    suite,
		withWrappers:             nil, // Set bellow.
		me:                       me,
		mySK:                     mySK,
		nodeIDs:                  nodeIDs,
		nodePKs:                  nodePKs,
		f:                        f,
		longTermSecretShare:      longTermSecretShare,
		dkg:                      nonce.New(suite, nodeIDs, nodePKs, f, me, mySK, log),
		dkgOutIndexes:            nil, // To be decided.
		dkgDecidedIndexProposals: nil, // To be received.
		dkgOutNonce:              nil, // To be decided.
		messageToSign:            nil, // Will be received later.
		dssPartialSigBuffer:      map[gpa.NodeID]*dss.PartialSig{},
		dssSigner:                nil, // Will be created when indexProposals and message to sign will be created.
		log:                      log,
	}
	d.msgWrapper = gpa.NewMsgWrapper(msgTypeWrapped, d.msgWrapperFunc)
	d.withWrappers = gpa.NewOwnHandler(me, d)
	return d
}

// DSS Specific Interface: Get a GPA instance to pass messages with all the intermediate layers.
func (d *dssImpl) AsGPA() gpa.GPA {
	return d.withWrappers
}

// Handle the input to the protocol.
func (d *dssImpl) Input(input gpa.Input) gpa.OutMessages {
	d.log.Debugf("Input %+v", input)
	switch input := input.(type) {
	case *inputStart:
		msgs := d.msgWrapper.WrapMessages(subsystemDKG, 0, d.dkg.Input(nonce.NewInputStart()))
		return d.tryHandleDkgOutput(msgs)
	case *inputDecided:
		return d.handleDecided(input)
	}
	panic(fmt.Errorf("unexpected input: %T: %+v", input, input))
}

// Handle the messages.
func (d *dssImpl) Message(msg gpa.Message) gpa.OutMessages {
	switch msgT := msg.(type) {
	case *msgPartialSig:
		d.log.Debugf("Message %+v", msg)
		return d.handlePartialSig(msgT)
	case *gpa.WrappingMsg:
		if msgT.Subsystem() == subsystemDKG && msgT.Index() == 0 {
			msgs := d.msgWrapper.WrapMessages(subsystemDKG, 0, d.dkg.Message(msgT.Wrapped()))
			return d.tryHandleDkgOutput(msgs)
		}
		panic(fmt.Errorf("unknown wrapped message %+v, wrapped %T: %v", msgT, msgT.Wrapped(), msgT.Wrapped()))
	default:
		panic(fmt.Errorf("unknown message %T: %v", msg, msg))
	}
}

// Provide the output, if any.
func (d *dssImpl) Output() gpa.Output {
	if d.dkgOutIndexes == nil && d.signature == nil {
		return nil
	}
	return &Output{
		ProposedIndexes: d.dkgOutIndexes,
		Signature:       d.signature,
	}
}

//nolint:gocyclo
func (d *dssImpl) tryHandleDkgOutput(msgs gpa.OutMessages) gpa.OutMessages {
	dkgOut := d.dkg.Output()
	if d.dkgOutIndexes == nil && dkgOut != nil && dkgOut.(*nonce.Output).Indexes != nil {
		d.dkgOutIndexes = dkgOut.(*nonce.Output).Indexes
	}
	if d.dkgOutNonce == nil && dkgOut != nil && dkgOut.(*nonce.Output).PriShare != nil {
		d.dkgOutNonce = tcrypto.NewDistKeyShare(
			dkgOut.(*nonce.Output).PriShare,
			dkgOut.(*nonce.Output).Commits,
			len(d.nodeIDs),
			dkgOut.(*nonce.Output).Threshold,
		)
		//
		// Create a partial signature.
		dssSigner, err := dss.NewDSS(d.suite, d.mySK, d.nodePKArray(), d.longTermSecretShare, d.dkgOutNonce, d.messageToSign, d.longTermSecretShare.Threshold())
		if err != nil {
			d.log.Error("Failed to create DSS Signer: %v", err)
			return msgs
		}
		d.dssSigner = dssSigner
		partialSig, err := d.dssSigner.PartialSig()
		if err != nil {
			d.log.Errorf("cannot create a partial signature: %v", err)
			return msgs
		}
		//
		// Process early sent partial signatures, if any.
		if len(d.dssPartialSigBuffer) > 0 {
			for nid := range d.dssPartialSigBuffer {
				err := d.dssSigner.ProcessPartialSig(d.dssPartialSigBuffer[nid])
				if err != nil {
					d.log.Warnf("Failed to process a buffered partial signature: %v", err)
				}
				delete(d.dssPartialSigBuffer, nid)
			}
		}
		//
		// Broadcast it (except the current node).
		for i := range d.nodeIDs {
			if d.nodeIDs[i] == d.me {
				continue
			}
			msgs.Add(&msgPartialSig{
				suite:      d.suite,
				sender:     d.me,
				recipient:  d.nodeIDs[i],
				partialSig: partialSig,
			})
		}
		//
		// Maybe we have everything for the signature already?
		if d.dssSigner.EnoughPartialSig() {
			sig, err := d.dssSigner.Signature()
			if err != nil {
				d.log.Errorf("unable to aggregate the signature: %v", err)
				return msgs
			}
			d.signature = sig
		}
	}
	return msgs
}

func (d *dssImpl) handlePartialSig(msg *msgPartialSig) gpa.OutMessages {
	if d.signature != nil {
		// Signature already aggregated, ignore the remaining shares.
		return nil
	}
	if d.dssSigner == nil {
		if _, ok := d.dssPartialSigBuffer[msg.sender]; ok {
			d.log.Warn("duplicate partial signature from %v", msg.sender)
			return nil
		}
		d.dssPartialSigBuffer[msg.sender] = msg.partialSig
		return nil
	}
	//
	// Then process the one received with the current message.
	err := d.dssSigner.ProcessPartialSig(msg.partialSig)
	if err != nil {
		d.log.Warnf("Failed to process a partial signature: %v", err)
		return nil
	}
	if !d.dssSigner.EnoughPartialSig() {
		return nil
	}

	sig, err := d.dssSigner.Signature()
	if err != nil {
		d.log.Errorf("unable to aggregate the signature: %v", err)
		return nil
	}
	d.signature = sig
	return nil
}

func (d *dssImpl) handleDecided(input *inputDecided) gpa.OutMessages {
	if d.dkgDecidedIndexProposals != nil {
		d.log.Warn("Duplicate will be dropped: DecidedIndexes=%+v", input.decidedIndexProposals)
		return nil
	}
	d.dkgDecidedIndexProposals = input.decidedIndexProposals
	d.messageToSign = input.messageToSign

	decisionInput := nonce.NewInputAgreementResult(input.decidedIndexProposals)
	msgs := d.msgWrapper.WrapMessages(subsystemDKG, 0, d.dkg.Input(decisionInput))
	return d.tryHandleDkgOutput(msgs)
}

func (d *dssImpl) nodePKArray() []kyber.Point {
	res := make([]kyber.Point, len(d.nodeIDs))
	for i := range res {
		res[i] = d.nodePKs[d.nodeIDs[i]]
	}
	return res
}

func (d *dssImpl) StatusString() string {
	return fmt.Sprintf("{DSS, dkg=%v}", d.dkg.StatusString())
}
