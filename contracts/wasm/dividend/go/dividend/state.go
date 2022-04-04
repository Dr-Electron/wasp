// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ArrayOfImmutableAddress struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableAddress) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableAddress) GetAddress(index uint32) wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(a.proxy.Index(index))
}

type MapAddressToImmutableUint64 struct {
	proxy wasmtypes.Proxy
}

func (m MapAddressToImmutableUint64) GetUint64(key wasmtypes.ScAddress) wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(m.proxy.Key(wasmtypes.AddressToBytes(key)))
}

type ImmutableDividendState struct {
	proxy wasmtypes.Proxy
}

// array with all the recipients of this dividend
func (s ImmutableDividendState) MemberList() ArrayOfImmutableAddress {
	return ArrayOfImmutableAddress{proxy: s.proxy.Root(StateMemberList)}
}

// map with all the recipient factors of this dividend
// owner of contract, the only one who can call 'member' func
// sum of all recipient factors
func (s ImmutableDividendState) Members() MapAddressToImmutableUint64 {
	return MapAddressToImmutableUint64{proxy: s.proxy.Root(StateMembers)}
}

func (s ImmutableDividendState) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateOwner))
}

func (s ImmutableDividendState) TotalFactor() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateTotalFactor))
}

type ArrayOfMutableAddress struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableAddress) AppendAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(a.proxy.Append())
}

func (a ArrayOfMutableAddress) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableAddress) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableAddress) GetAddress(index uint32) wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(a.proxy.Index(index))
}

type MapAddressToMutableUint64 struct {
	proxy wasmtypes.Proxy
}

func (m MapAddressToMutableUint64) Clear() {
	m.proxy.ClearMap()
}

func (m MapAddressToMutableUint64) GetUint64(key wasmtypes.ScAddress) wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(m.proxy.Key(wasmtypes.AddressToBytes(key)))
}

type MutableDividendState struct {
	proxy wasmtypes.Proxy
}

func (s MutableDividendState) AsImmutable() ImmutableDividendState {
	return ImmutableDividendState(s)
}

// array with all the recipients of this dividend
func (s MutableDividendState) MemberList() ArrayOfMutableAddress {
	return ArrayOfMutableAddress{proxy: s.proxy.Root(StateMemberList)}
}

// map with all the recipient factors of this dividend
// owner of contract, the only one who can call 'member' func
// sum of all recipient factors
func (s MutableDividendState) Members() MapAddressToMutableUint64 {
	return MapAddressToMutableUint64{proxy: s.proxy.Root(StateMembers)}
}

func (s MutableDividendState) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateOwner))
}

func (s MutableDividendState) TotalFactor() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateTotalFactor))
}
