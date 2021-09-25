// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct CallOnChainCall {
    pub func:    ScFunc,
    pub params:  MutableCallOnChainParams,
    pub results: ImmutableCallOnChainResults,
}

pub struct CheckContextFromFullEPCall {
    pub func:   ScFunc,
    pub params: MutableCheckContextFromFullEPParams,
}

pub struct DoNothingCall {
    pub func: ScFunc,
}

pub struct GetMintedSupplyCall {
    pub func:    ScFunc,
    pub results: ImmutableGetMintedSupplyResults,
}

pub struct IncCounterCall {
    pub func: ScFunc,
}

pub struct InitCall {
    pub func: ScInitFunc,
}

pub struct PassTypesFullCall {
    pub func:   ScFunc,
    pub params: MutablePassTypesFullParams,
}

pub struct RunRecursionCall {
    pub func:    ScFunc,
    pub params:  MutableRunRecursionParams,
    pub results: ImmutableRunRecursionResults,
}

pub struct SendToAddressCall {
    pub func:   ScFunc,
    pub params: MutableSendToAddressParams,
}

pub struct SetIntCall {
    pub func:   ScFunc,
    pub params: MutableSetIntParams,
}

pub struct TestBlockContext1Call {
    pub func: ScFunc,
}

pub struct TestBlockContext2Call {
    pub func: ScFunc,
}

pub struct TestCallPanicFullEPCall {
    pub func: ScFunc,
}

pub struct TestCallPanicViewEPFromFullCall {
    pub func: ScFunc,
}

pub struct TestChainOwnerIDFullCall {
    pub func:    ScFunc,
    pub results: ImmutableTestChainOwnerIDFullResults,
}

pub struct TestEventLogDeployCall {
    pub func: ScFunc,
}

pub struct TestEventLogEventDataCall {
    pub func: ScFunc,
}

pub struct TestEventLogGenericDataCall {
    pub func:   ScFunc,
    pub params: MutableTestEventLogGenericDataParams,
}

pub struct TestPanicFullEPCall {
    pub func: ScFunc,
}

pub struct WithdrawToChainCall {
    pub func:   ScFunc,
    pub params: MutableWithdrawToChainParams,
}

pub struct CheckContextFromViewEPCall {
    pub func:   ScView,
    pub params: MutableCheckContextFromViewEPParams,
}

pub struct FibonacciCall {
    pub func:    ScView,
    pub params:  MutableFibonacciParams,
    pub results: ImmutableFibonacciResults,
}

pub struct GetCounterCall {
    pub func:    ScView,
    pub results: ImmutableGetCounterResults,
}

pub struct GetIntCall {
    pub func:    ScView,
    pub params:  MutableGetIntParams,
    pub results: ImmutableGetIntResults,
}

pub struct GetStringValueCall {
    pub func:    ScView,
    pub params:  MutableGetStringValueParams,
    pub results: ImmutableGetStringValueResults,
}

pub struct JustViewCall {
    pub func: ScView,
}

pub struct PassTypesViewCall {
    pub func:   ScView,
    pub params: MutablePassTypesViewParams,
}

pub struct TestCallPanicViewEPFromViewCall {
    pub func: ScView,
}

pub struct TestChainOwnerIDViewCall {
    pub func:    ScView,
    pub results: ImmutableTestChainOwnerIDViewResults,
}

pub struct TestPanicViewEPCall {
    pub func: ScView,
}

pub struct TestSandboxCallCall {
    pub func:    ScView,
    pub results: ImmutableTestSandboxCallResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn call_on_chain(_ctx: & dyn ScFuncCallContext) -> CallOnChainCall {
        let mut f = CallOnChainCall {
            func:    ScFunc::new(HSC_NAME, HFUNC_CALL_ON_CHAIN),
            params:  MutableCallOnChainParams { id: 0 },
            results: ImmutableCallOnChainResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn check_context_from_full_ep(_ctx: & dyn ScFuncCallContext) -> CheckContextFromFullEPCall {
        let mut f = CheckContextFromFullEPCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_CHECK_CONTEXT_FROM_FULL_EP),
            params: MutableCheckContextFromFullEPParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn do_nothing(_ctx: & dyn ScFuncCallContext) -> DoNothingCall {
        DoNothingCall {
            func: ScFunc::new(HSC_NAME, HFUNC_DO_NOTHING),
        }
    }
    pub fn get_minted_supply(_ctx: & dyn ScFuncCallContext) -> GetMintedSupplyCall {
        let mut f = GetMintedSupplyCall {
            func:    ScFunc::new(HSC_NAME, HFUNC_GET_MINTED_SUPPLY),
            results: ImmutableGetMintedSupplyResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn inc_counter(_ctx: & dyn ScFuncCallContext) -> IncCounterCall {
        IncCounterCall {
            func: ScFunc::new(HSC_NAME, HFUNC_INC_COUNTER),
        }
    }
    pub fn init(_ctx: & dyn ScFuncCallContext) -> InitCall {
        InitCall {
            func: ScInitFunc::new(HSC_NAME, HFUNC_INIT),
        }
    }
    pub fn pass_types_full(_ctx: & dyn ScFuncCallContext) -> PassTypesFullCall {
        let mut f = PassTypesFullCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_PASS_TYPES_FULL),
            params: MutablePassTypesFullParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn run_recursion(_ctx: & dyn ScFuncCallContext) -> RunRecursionCall {
        let mut f = RunRecursionCall {
            func:    ScFunc::new(HSC_NAME, HFUNC_RUN_RECURSION),
            params:  MutableRunRecursionParams { id: 0 },
            results: ImmutableRunRecursionResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn send_to_address(_ctx: & dyn ScFuncCallContext) -> SendToAddressCall {
        let mut f = SendToAddressCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SEND_TO_ADDRESS),
            params: MutableSendToAddressParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn set_int(_ctx: & dyn ScFuncCallContext) -> SetIntCall {
        let mut f = SetIntCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_SET_INT),
            params: MutableSetIntParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn test_block_context1(_ctx: & dyn ScFuncCallContext) -> TestBlockContext1Call {
        TestBlockContext1Call {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_BLOCK_CONTEXT1),
        }
    }
    pub fn test_block_context2(_ctx: & dyn ScFuncCallContext) -> TestBlockContext2Call {
        TestBlockContext2Call {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_BLOCK_CONTEXT2),
        }
    }
    pub fn test_call_panic_full_ep(_ctx: & dyn ScFuncCallContext) -> TestCallPanicFullEPCall {
        TestCallPanicFullEPCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_CALL_PANIC_FULL_EP),
        }
    }
    pub fn test_call_panic_view_ep_from_full(_ctx: & dyn ScFuncCallContext) -> TestCallPanicViewEPFromFullCall {
        TestCallPanicViewEPFromFullCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_CALL_PANIC_VIEW_EP_FROM_FULL),
        }
    }
    pub fn test_chain_owner_id_full(_ctx: & dyn ScFuncCallContext) -> TestChainOwnerIDFullCall {
        let mut f = TestChainOwnerIDFullCall {
            func:    ScFunc::new(HSC_NAME, HFUNC_TEST_CHAIN_OWNER_ID_FULL),
            results: ImmutableTestChainOwnerIDFullResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn test_event_log_deploy(_ctx: & dyn ScFuncCallContext) -> TestEventLogDeployCall {
        TestEventLogDeployCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_EVENT_LOG_DEPLOY),
        }
    }
    pub fn test_event_log_event_data(_ctx: & dyn ScFuncCallContext) -> TestEventLogEventDataCall {
        TestEventLogEventDataCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_EVENT_LOG_EVENT_DATA),
        }
    }
    pub fn test_event_log_generic_data(_ctx: & dyn ScFuncCallContext) -> TestEventLogGenericDataCall {
        let mut f = TestEventLogGenericDataCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_TEST_EVENT_LOG_GENERIC_DATA),
            params: MutableTestEventLogGenericDataParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn test_panic_full_ep(_ctx: & dyn ScFuncCallContext) -> TestPanicFullEPCall {
        TestPanicFullEPCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_PANIC_FULL_EP),
        }
    }
    pub fn withdraw_to_chain(_ctx: & dyn ScFuncCallContext) -> WithdrawToChainCall {
        let mut f = WithdrawToChainCall {
            func:   ScFunc::new(HSC_NAME, HFUNC_WITHDRAW_TO_CHAIN),
            params: MutableWithdrawToChainParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn check_context_from_view_ep(_ctx: & dyn ScViewCallContext) -> CheckContextFromViewEPCall {
        let mut f = CheckContextFromViewEPCall {
            func:   ScView::new(HSC_NAME, HVIEW_CHECK_CONTEXT_FROM_VIEW_EP),
            params: MutableCheckContextFromViewEPParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn fibonacci(_ctx: & dyn ScViewCallContext) -> FibonacciCall {
        let mut f = FibonacciCall {
            func:    ScView::new(HSC_NAME, HVIEW_FIBONACCI),
            params:  MutableFibonacciParams { id: 0 },
            results: ImmutableFibonacciResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_counter(_ctx: & dyn ScViewCallContext) -> GetCounterCall {
        let mut f = GetCounterCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_COUNTER),
            results: ImmutableGetCounterResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn get_int(_ctx: & dyn ScViewCallContext) -> GetIntCall {
        let mut f = GetIntCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_INT),
            params:  MutableGetIntParams { id: 0 },
            results: ImmutableGetIntResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn get_string_value(_ctx: & dyn ScViewCallContext) -> GetStringValueCall {
        let mut f = GetStringValueCall {
            func:    ScView::new(HSC_NAME, HVIEW_GET_STRING_VALUE),
            params:  MutableGetStringValueParams { id: 0 },
            results: ImmutableGetStringValueResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
    pub fn just_view(_ctx: & dyn ScViewCallContext) -> JustViewCall {
        JustViewCall {
            func: ScView::new(HSC_NAME, HVIEW_JUST_VIEW),
        }
    }
    pub fn pass_types_view(_ctx: & dyn ScViewCallContext) -> PassTypesViewCall {
        let mut f = PassTypesViewCall {
            func:   ScView::new(HSC_NAME, HVIEW_PASS_TYPES_VIEW),
            params: MutablePassTypesViewParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn test_call_panic_view_ep_from_view(_ctx: & dyn ScViewCallContext) -> TestCallPanicViewEPFromViewCall {
        TestCallPanicViewEPFromViewCall {
            func: ScView::new(HSC_NAME, HVIEW_TEST_CALL_PANIC_VIEW_EP_FROM_VIEW),
        }
    }
    pub fn test_chain_owner_id_view(_ctx: & dyn ScViewCallContext) -> TestChainOwnerIDViewCall {
        let mut f = TestChainOwnerIDViewCall {
            func:    ScView::new(HSC_NAME, HVIEW_TEST_CHAIN_OWNER_ID_VIEW),
            results: ImmutableTestChainOwnerIDViewResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
    pub fn test_panic_view_ep(_ctx: & dyn ScViewCallContext) -> TestPanicViewEPCall {
        TestPanicViewEPCall {
            func: ScView::new(HSC_NAME, HVIEW_TEST_PANIC_VIEW_EP),
        }
    }
    pub fn test_sandbox_call(_ctx: & dyn ScViewCallContext) -> TestSandboxCallCall {
        let mut f = TestSandboxCallCall {
            func:    ScView::new(HSC_NAME, HVIEW_TEST_SANDBOX_CALL),
            results: ImmutableTestSandboxCallResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
}

// @formatter:on
