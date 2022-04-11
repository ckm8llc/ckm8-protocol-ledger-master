# ckm8 Blockchain Ledger Protocol

The ckm8 Blockchain Ledger is a Proof-of-Stake decentralized ledger designed for the video streaming industry. It powers the ckm8 token economy which incentives end users to share their redundant bandwidth and storage resources, and encourage them to engage more actively with video platforms and content creators. The ledger employs a novel [multi-level BFT consensus engine], which supports high transaction throughput, fast block confirmation, and allows mass participation in the consensus process. Off-chain payment support is built directly into the ledger through the resource-oriented micropayment pool, which is designed specifically to achieve the “pay-per-byte” granularity for streaming use cases. Moreover, the ledger storage system leverages the microservice architecture and reference counting based history pruning techniques, and is thus able to adapt to different computing environments, ranging from high-end data center server clusters to commodity PCs and laptops. The ledger also supports Turing-Complete smart contracts, which enables rich user experiences for DApps built on top of the ckm8 Ledger. For more technical details, please refer to our [technical whitepaper](docs/ckm8-technical-whitepaper.pdf) and [2019 IEEE ICBC paper](https://arxiv.org/pdf/1911.04698.pdf) "Scalable BFT Consensus Mechanism Through Aggregated
Signature Gossip".

To learn more about the ckm8 Network in general, please visit the **ckm8 Documentation site**: https://docs.ckm8token.org/docs/what-is-ckm8-network.

## Table of Contents
- [Setup](#setup)
- [Smart Contract and DApp Development on ckm8](#smart-contract-and-dapp-development-on-ckm8)

## Setup

### Intall Go

Install Go and set environment variables `GOPATH` , `GOBIN`, and `PATH`. The current code base should compile with **Go 1.14.2**. On macOS, install Go with the following command

```
brew install go@1.14.1
brew link go@1.14.1 --force
```

### Build and Install

Next, clone this repo into your `$GOPATH`. The path should look like this: `$GOPATH/src/github.com/ckm8token/ckm8`

```
git clonehttps://github.com/fsmile2/ckm8/ckm8-protocol-ledger.git $GOPATH/src/github.com/ckm8token/ckm8
export ckm8_HOME=$GOPATH/src/github.com/ckm8token/ckm8
cd $ckm8_HOME
```

Now, execute the following commands to build the ckm8 binaries under `$GOPATH/bin`. Two binaries `ckm8` and `ckm8cli` are generated. `ckm8` can be regarded as the launcher of the ckm8 Ledger node, and `ckm8cli` is a wallet with command line tools to interact with the ledger.

```
export GO111MODULE=on
make install
```

#### Notes for Linux binary compilation
The build and install process on **Linux** is similar, but note that Ubuntu 18.04.4 LTS / Centos 8 or higher version is required for the compilation. 

#### Notes for Windows binary compilation
The Windows binary can be cross-compiled from macOS. To cross-compile a **Windows** binary, first make sure `mingw64` is installed (`brew install mingw-w64`) on your macOS. Then you can cross-compile the Windows binary with the following command:

```
make exe
```

You'll also need to place three `.dll` files `libgcc_s_seh-1.dll`, `libstdc++-6.dll`, `libwinpthread-1.dll` under the same folder as `ckm8.exe` and `ckm8cli.exe`.


### Run Unit Tests
Run unit tests with the command below
```
make test_unit
```

## Smart Contract and DApp Development on ckm8

ckm8 provides full support for Turing-Complete smart contract, and is EVM compatible. To start developing on the ckm8 Blockchain, please check out the following links:

### Smart Contracts
* Smart contract and DApp development Overview: [link here](https://docs.ckm8token.org/docs/turing-complete-smart-contract-support). 
* Tutorials on how to interact with the ckm8 blockchain through [Metamask](https://docs.ckm8token.org/docs/web3-stack-metamask), [Truffle](https://docs.ckm8token.org/docs/web3-stack-truffle), [Hardhat](https://docs.ckm8token.org/docs/web3-stack-hardhat), [web3.js](https://docs.ckm8token.org/docs/web3-stack-web3js), and [ethers.js](https://docs.ckm8token.org/docs/web3-stack-hardhat).
* TNT20 Token (i.e. ERC20 on ckm8) integration guide: [link here](https://docs.ckm8token.org/docs/ckm8-blockchain-tnt20-token-integration-guide).

### Local Test Environment Setup
* Launching a local privatenet: [link here](https://docs.ckm8token.org/docs/launch-a-local-privatenet).
* Command line tools: [link here](https://docs.ckm8token.org/docs/command-line-tool).
* Connect to the [Testnet](https://docs.ckm8token.org/docs/connect-to-the-testnet), and the [Mainnet](https://docs.ckm8token.org/docs/connect-to-the-mainnet).
* Node configuration: [link here](https://docs.ckm8token.org/docs/ckm8-blockchain-node-configuration).

### API References
* Native RPC API references: [link here](https://docs.ckm8token.org/docs/rpc-api-reference).
* Ethereum RPC API support: [link here](https://docs.ckm8token.org/docs/web3-stack-eth-rpc-support).

