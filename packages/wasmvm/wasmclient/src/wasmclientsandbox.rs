// // Copyright 2020 IOTA Stiftung
// // SPDX-License-Identifier: Apache-2.0

use crate::*;
use wasmlib::*;

pub struct WasmClientSandbox(WasmClientContext);

pub trait ScHost {
    fn export_name(&self, index: i32, name: &str);
    fn sandbox(self, func_num: i32, params: &[u8]) -> Vec<u8>;
    fn state_delete(&self, key: &[u8]);
    fn state_exists(&self, key: &[u8]) -> bool;
    fn state_get(&self, key: &[u8]) -> Vec<u8>;
    fn state_set(&self, key: &[u8], value: &[u8]);
}

impl ScHost for WasmClientSandbox {
    fn export_name(&self, _index: i32, _name: &str) {
        panic!("WasmClientContext.ExportName")
    }

    fn sandbox(self, func_num: i32, args: &[u8]) -> Vec<u8> {
        match func_num {
            wasmlib::FN_CALL => return self.fn_call(args).unwrap(),
            wasmlib::FN_POST => return self.fn_post(args).unwrap(),
            wasmlib::FN_UTILS_BECH32_DECODE => return self.fn_utils_bech32_decode(args).unwrap(),
            wasmlib::FN_UTILS_BECH32_ENCODE => return self.fn_utils_bech32_encode(args).unwrap(),
            wasmlib::FN_UTILS_HASH_NAME => return self.fn_utils_hash_name(args).unwrap(),
            _ => panic!("implement WasmClientContext.Sandbox"),
        }
    }

    fn state_delete(&self, _key: &[u8]) {
        panic!("WasmClientContext.StateDelete")
    }

    fn state_exists(&self, _key: &[u8]) -> bool {
        panic!("WasmClientContext.StateExists")
    }

    fn state_get(&self, _key: &[u8]) -> Vec<u8> {
        panic!("WasmClientContext.StateGet")
    }

    fn state_set(&self, _key: &[u8], _value: &[u8]) {
        panic!("WasmClientContext.StateSet")
    }
}

impl WasmClientSandbox {
    pub fn new() -> Self {
        return WasmClientSandbox(WasmClientContext::new());
    }
    pub fn fn_call(mut self, args: &[u8]) -> Result<Vec<u8>, String> {
        let req = wasmrequests::PostRequest::from_bytes(args);
        if req.contract == self.0.sc_hname {
            self.0.err = String::from(format!("unknown contract: {}", req.contract.to_string()));
            return Err(self.0.err);
        }

        return self.0.svc_client.call_view_by_hname(
            self.0.chain_id,
            req.contract,
            req.function,
            &req.params,
        );
    }

    pub fn fn_post(mut self, args: &[u8]) -> Result<Vec<u8>, String> {
        let req = wasmrequests::PostRequest::from_bytes(args);
        if req.chain_id == self.0.chain_id {
            self.0.err = String::from(format!("unknown contract: {}", req.contract.to_string()));
            return Err(self.0.err);
        }
        if req.contract == self.0.sc_hname {
            self.0.err = String::from(format!("unknown contract: {}", req.contract.to_string()));
            return Err(self.0.err);
        }
        let sc_assets = wasmlib::ScAssets::new(&req.transfer);
        self.0.svc_client.post_request(
            self.0.chain_id,
            req.contract,
            req.function,
            &req.params,
            sc_assets,
            self.0.key_pair,
        )?;
        return Ok(Vec::new());
    }

    pub fn fn_utils_bech32_decode(&self, args: &[u8]) -> Result<Vec<u8>, String> {
        let bech32 = wasmlib::string_from_bytes(args);
        let addr = codec::bech32_decode(&bech32)?;
        let cvt = wasmconvertor::WasmConvertor::new();
        return Ok(cvt.sc_address(addr).to_bytes());
    }

    pub fn fn_utils_bech32_encode(&self, args: &[u8]) -> Result<Vec<u8>, String> {
        let cvt = wasmconvertor::WasmConvertor::new();
        let sc_address = wasmtypes::address_from_bytes(args);
        let addr = cvt.isc_address(&sc_address);
        let bech32 = codec::bech32_encode(addr);
        return Ok(bech32.into_bytes());
    }

    pub fn fn_utils_hash_name(&self, args: &[u8]) -> Result<Vec<u8>, String> {
        let s = match std::str::from_utf8(args) {
            Ok(v) => v,
            Err(e) => return Err(String::from(format!("invalid hname: {}", e))),
        };
        return Ok(wasmtypes::hname_from_string(s).to_bytes());
    }
}
