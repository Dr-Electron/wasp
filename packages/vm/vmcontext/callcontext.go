package vmcontext

import (
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/core/accounts/commonaccount"
)

// pushCallContextWithMoveAssets moves funds from the caller to the non-core target before pushing new context to the stack
func (vmctx *VMContext) pushCallContextWithMoveAssets(targetContract iscp.Hname, params dict.Dict, transfer *iscp.Assets) {
	if transfer != nil && commonaccount.IsCoreHname(targetContract) {
		// if target contract is one of core contracts, transfer is ignored
		// this makes it impossible to send funds to 'common account' with ordinary smart contract call
		// transfer will end up either in the non-core target account or in the sender's account
		transfer = nil
		vmctx.Debugf("transfer ignored for core targetContract")
	}
	if transfer != nil {
		targetAccount := iscp.NewAgentID(vmctx.ChainID().AsAddress(), targetContract)
		var sourceAccount *iscp.AgentID
		if len(vmctx.callStack) == 0 {
			sourceAccount = vmctx.req.SenderAccount()
		} else {
			sourceAccount = vmctx.AccountID()
		}
		vmctx.mustMoveBetweenAccounts(sourceAccount, targetAccount, transfer)
	}
	vmctx.pushCallContext(targetContract, params, transfer)
}

const traceStack = false

func (vmctx *VMContext) getCaller() *iscp.AgentID {
	if len(vmctx.callStack) > 0 {
		return vmctx.MyAgentID()
	}
	if vmctx.req == nil {
		// core call (e.g. saving the anchor ID)
		return vmctx.chainOwnerID
	}
	// request context
	return vmctx.req.SenderAccount()
}

func (vmctx *VMContext) pushCallContext(contract iscp.Hname, params dict.Dict, transfer *iscp.Assets) {
	ctx := &callContext{
		caller:   vmctx.getCaller(),
		contract: contract,
		params:   params.Clone(),
		transfer: transfer,
	}
	if traceStack {
		vmctx.Debugf("+++++++++++ PUSH %d, stack depth = %d caller = %s", contract, len(vmctx.callStack), ctx.caller)
	}
	vmctx.callStack = append(vmctx.callStack, ctx)
}

func (vmctx *VMContext) popCallContext() {
	if traceStack {
		vmctx.Debugf("+++++++++++ POP @ depth %d", len(vmctx.callStack))
	}
	vmctx.callStack[len(vmctx.callStack)-1] = nil // for GC
	vmctx.callStack = vmctx.callStack[:len(vmctx.callStack)-1]
}

func (vmctx *VMContext) getCallContext() *callContext {
	if len(vmctx.callStack) == 0 {
		panic("getCallContext: stack is empty")
	}
	return vmctx.callStack[len(vmctx.callStack)-1]
}

func (vmctx *VMContext) callCore(c *coreutil.ContractInfo, f func(s kv.KVStore)) {
	vmctx.pushCallContext(c.Hname(), nil, nil)
	defer vmctx.popCallContext()

	f(vmctx.State())
}
