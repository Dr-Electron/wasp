(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[2699],{4009:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return i},contentTitle:function(){return s},metadata:function(){return l},toc:function(){return p},default:function(){return d}});var n=r(2122),o=r(9756),c=(r(7294),r(3905)),a=["components"],i={},s="Core Contracts",l={unversionedId:"guide/core_concepts/core_contracts/overview",id:"guide/core_concepts/core_contracts/overview",isDocsHomePage:!1,title:"Core Contracts",description:"Deploying a chain automatically means deployment of all core smart contracts on",source:"@site/docs/guide/core_concepts/core_contracts/overview.md",sourceDirName:"guide/core_concepts/core_contracts",slug:"/guide/core_concepts/core_contracts/overview",permalink:"/docs/guide/core_concepts/core_contracts/overview",editUrl:"https://github.com/iotaledger/chronicle.rs/tree/main/docs/docs/guide/core_concepts/core_contracts/overview.md",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"sandbox",permalink:"/docs/guide/core_concepts/sandbox"},next:{title:"root",permalink:"/docs/guide/core_concepts/core_contracts/root"}},p=[],u={toc:p};function d(e){var t=e.components,r=(0,o.Z)(e,a);return(0,c.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,c.kt)("h1",{id:"core-contracts"},"Core Contracts"),(0,c.kt)("p",null,"Deploying a chain automatically means deployment of all core smart contracts on\nit. The core contracts are responsible for the vital functions of the chain and\nprovide infrastructure for all other smart contracts:"),(0,c.kt)("ul",null,(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("p",{parentName:"li"},(0,c.kt)("a",{parentName:"p",href:"/docs/guide/core_concepts/core_contracts/root"},(0,c.kt)("strong",{parentName:"a"},"root"))," - Responsible for the initialization of the chain, maintains\nthe global parameters, and the registry of deployed contracts. It also handles\nfees and performs other functions.")),(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("p",{parentName:"li"},(0,c.kt)("a",{parentName:"p",href:"/docs/guide/core_concepts/core_contracts/default"},(0,c.kt)("strong",{parentName:"a"},"_default"))," - Any request that cannot be handled by any of the\nother deployed contracts ends up here.")),(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("p",{parentName:"li"},(0,c.kt)("a",{parentName:"p",href:"/docs/guide/core_concepts/core_contracts/accounts"},(0,c.kt)("strong",{parentName:"a"},"accounts"))," - Responsible for the on-chain ledger of accounts. The\non-chain accounts contain colored tokens, which are controlled by smart\ncontracts and addresses on the UTXO Ledger.")),(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("p",{parentName:"li"},(0,c.kt)("a",{parentName:"p",href:"/docs/guide/core_concepts/core_contracts/blob"},(0,c.kt)("strong",{parentName:"a"},"blob"))," - Responsible for the immutable registry of binary objects of\narbitrary size. One blob is a collection of named binary chunks of data. For\nexample, a blob can be used to store a collections of ",(0,c.kt)("em",{parentName:"p"},"wasm")," binaries, needed\nto deploy ",(0,c.kt)("em",{parentName:"p"},"WebAssembly")," smart contracts. Each blob in the registry is\nreferenced by its hash which is deterministically calculated from its data.")),(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("p",{parentName:"li"},(0,c.kt)("a",{parentName:"p",href:"/docs/guide/core_concepts/core_contracts/blocklog"},(0,c.kt)("strong",{parentName:"a"},"blocklog"))," - Keeps track of the blocks, requests and event that were\nprocessed by the chain."))))}d.isMDXComponent=!0},3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return p},kt:function(){return f}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var s=n.createContext({}),l=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},p=function(e){var t=l(e.components);return n.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,c=e.originalType,s=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),d=l(r),f=o,m=d["".concat(s,".").concat(f)]||d[f]||u[f]||c;return r?n.createElement(m,a(a({ref:t},p),{},{components:r})):n.createElement(m,a({ref:t},p))}));function f(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var c=r.length,a=new Array(c);a[0]=d;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:o,a[1]=i;for(var l=2;l<c;l++)a[l]=r[l];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"}}]);