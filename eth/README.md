# Jinx Ethereum

Welcome to Jinx Ethereum, a modular framework for injecting a Go-Ethereum (geth) EVM into any 
underlying consensus layer. This folder's directory structure closely resembles that of geth, as it
is meant to be a thin wrapper around the existing geth codebase. The following architecture diagram
shows how Jinx Ethereum integrates into the application level of a host chain.

![jinx_architecture.png](../docs/web/public/jinx_architecture.png)

## api

`api` includes the public Chain API that Jinx Ethereum exports.
 
## core

`core` includes the Jinx Core logic that runs the EVM: process blocks, transactions, and state
transitions. This encapsulates **State Processor** and **Embedded Host Chain** in the architecture
diagram.

## rpc

`rpc` includes rpc service that can be injected into the host chain's JSON-RPC server. This 
encapsulates **RPC Backend** in the architecture diagram. 

## [provider.go](https://github.com/berachain/jinx/blob/main/eth/provider.go) 

The `JinxProvider` can be exported and used by the host chain.
