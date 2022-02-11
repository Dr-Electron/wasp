package blocklog

import (
	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/wasp/packages/vm/core/errors"
	"github.com/iotaledger/wasp/packages/vm/core/errors/commonerrors"
	error2 "github.com/iotaledger/wasp/packages/vm/vmerrors"
	"testing"
	"time"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/stretchr/testify/require"
)

var FailedToLoadError *error2.ErrorDefinition = commonerrors.RegisterGlobalError(1, "Failed to load %v, with %v on %v")

func TestSimpleErrorSerialization(t *testing.T) {
	mu := marshalutil.New()

	// Initial error
	blockError := FailedToLoadError.Create("placeBet", "destroy", "setAdmin")

	// Serialize error to bytes
	err := blockError.Serialize(mu)
	require.NoError(t, err)

	// Recreate error from bytes
	newError, err := errors.ErrorFromBytes(mu, nil)
	require.NoError(t, err)

	// Validate that properties are the same
	require.EqualValues(t, blockError.Hash(), newError.Hash())
	require.EqualValues(t, blockError.Params, newError.Params)
	require.EqualValues(t, blockError.MessageFormat, newError.MessageFormat)

	// Validate that error returns a proper error type
	require.Error(t, newError)

	t.Log(newError.Error())
}

func TestSerdeRequestReceipt(t *testing.T) {
	nonce := uint64(time.Now().UnixNano())
	req := iscp.NewOffLedgerRequest(iscp.RandomChainID(), iscp.Hn("0"), iscp.Hn("0"), nil, nonce)
	//testError, err := GeneralErrorCollection.Create(1, 1)

	rec := &RequestReceipt{
		Request: req,
		//Error:   testError,
	}
	forward := rec.Bytes()
	back, err := RequestReceiptFromBytes(forward)
	require.NoError(t, err)
	require.EqualValues(t, forward, back.Bytes())
}
