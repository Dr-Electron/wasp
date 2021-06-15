// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package native

import (
	"fmt"
	"sync"

	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/coreutil"
	"github.com/iotaledger/wasp/packages/hashing"
)

const VMType = "examplevm"

var (
	allExamples      = make(map[hashing.HashValue]coretypes.VMProcessor)
	allExamplesMutex = &sync.Mutex{}
)

// AddProcessor adds new processor to the runtime registry of example processors.
// The 'proc' represents executable of the specific smart contract.
// It must implement coretypes.VMProcessor
func AddProcessor(c *coreutil.ContractInterface) {
	allExamplesMutex.Lock()
	defer allExamplesMutex.Unlock()
	allExamples[c.ProgramHash] = c
	fmt.Printf("----- AddProcessor: name: '%s', program hash: %s, description: '%s'\n",
		c.Name, c.ProgramHash.String(), c.Description)
}

// GetProcessor retrieves smart contract processor (VM) by the hash (with existence flag)
func GetProcessor(progHash hashing.HashValue) (coretypes.VMProcessor, bool) {
	ret, ok := allExamples[progHash]
	if !ok {
		return nil, false
	}
	return ret, true
}
