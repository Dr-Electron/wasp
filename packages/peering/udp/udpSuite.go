package udp

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import (
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Suite interface {
	pairing.Suite
	kyber.Group
}
