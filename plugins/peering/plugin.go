package peering

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import (
	"github.com/iotaledger/hive.go/daemon"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/hive.go/node"
	"github.com/iotaledger/wasp/packages/parameters"
	peering_pkg "github.com/iotaledger/wasp/packages/peering"
	peering_udp "github.com/iotaledger/wasp/packages/peering/udp"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/labstack/gommon/log"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/key"
)

const (
	pluginName = "Peering"
)

var (
	defaultNetworkProvider *peering_udp.NetImpl // A singleton instance.
)

// Init is an entry point for this plugin.
func Init(suite *pairing.SuiteBn256) *node.Plugin {
	configure := func(_ *node.Plugin) {
		var err error
		var nodeKeyPair *key.Pair
		if nodeKeyPair, err = registry.DefaultRegistry().GetNodeIdentity(); err != nil {
			panic(err)
		}
		defaultNetworkProvider, err = peering_udp.NewNetworkProvider(
			parameters.GetString(parameters.PeeringMyNetId),
			parameters.GetInt(parameters.PeeringPort),
			nodeKeyPair,
			suite,
			logger.NewLogger(pluginName),
		)
		if err != nil {
			panic(err)
		}
		log.Infof(
			"--------------------------------- NetID is %s -----------------------------------",
			defaultNetworkProvider.Self().NetID(),
		)
	}
	run := func(_ *node.Plugin) {
		err := daemon.BackgroundWorker(
			"WaspPeering",
			defaultNetworkProvider.Run,
			parameters.PriorityPeering,
		)
		if err != nil {
			panic(err)
		}
	}
	return node.NewPlugin(pluginName, node.Enabled, configure, run)
}

// DefaultNetworkProvider returns the default network provider implementation.
func DefaultNetworkProvider() peering_pkg.NetworkProvider {
	return defaultNetworkProvider
}
