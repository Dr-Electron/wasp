// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use donatewithfeedback::*;
use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;
use crate::structs::*;

mod consts;
mod contract;
mod params;
mod results;
mod state;
mod structs;

mod donatewithfeedback;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_DONATE,
    	FUNC_WITHDRAW,
    	VIEW_DONATION,
    	VIEW_DONATION_INFO,
	],
    funcs: &[
    	func_donate_thunk,
    	func_withdraw_thunk,
	],
    views: &[
    	view_donation_thunk,
    	view_donation_info_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct DonateContext {
	params: ImmutableDonateParams,
	state: MutableDonateWithFeedbackState,
}

fn func_donate_thunk(ctx: &ScFuncContext) {
	ctx.log("donatewithfeedback.funcDonate");
	let f = DonateContext {
		params: ImmutableDonateParams { proxy: params_proxy() },
		state: MutableDonateWithFeedbackState { proxy: state_proxy() },
	};
	func_donate(ctx, &f);
	ctx.log("donatewithfeedback.funcDonate ok");
}

pub struct WithdrawContext {
	params: ImmutableWithdrawParams,
	state: MutableDonateWithFeedbackState,
}

fn func_withdraw_thunk(ctx: &ScFuncContext) {
	ctx.log("donatewithfeedback.funcWithdraw");
	let f = WithdrawContext {
		params: ImmutableWithdrawParams { proxy: params_proxy() },
		state: MutableDonateWithFeedbackState { proxy: state_proxy() },
	};
	ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

	func_withdraw(ctx, &f);
	ctx.log("donatewithfeedback.funcWithdraw ok");
}

pub struct DonationContext {
	params: ImmutableDonationParams,
	results: MutableDonationResults,
	state: ImmutableDonateWithFeedbackState,
}

fn view_donation_thunk(ctx: &ScViewContext) {
	ctx.log("donatewithfeedback.viewDonation");
	let f = DonationContext {
		params: ImmutableDonationParams { proxy: params_proxy() },
		results: MutableDonationResults { proxy: results_proxy() },
		state: ImmutableDonateWithFeedbackState { proxy: state_proxy() },
	};
	ctx.require(f.params.nr().exists(), "missing mandatory nr");
	view_donation(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("donatewithfeedback.viewDonation ok");
}

pub struct DonationInfoContext {
	results: MutableDonationInfoResults,
	state: ImmutableDonateWithFeedbackState,
}

fn view_donation_info_thunk(ctx: &ScViewContext) {
	ctx.log("donatewithfeedback.viewDonationInfo");
	let f = DonationInfoContext {
		results: MutableDonationInfoResults { proxy: results_proxy() },
		state: ImmutableDonateWithFeedbackState { proxy: state_proxy() },
	};
	view_donation_info(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("donatewithfeedback.viewDonationInfo ok");
}
