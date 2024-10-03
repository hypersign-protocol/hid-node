# Hypersign Identity Network

[![build-test](https://github.com/hypersign-protocol/hid-node/actions/workflows/build.yml/badge.svg)](https://github.com/hypersign-protocol/hid-node/actions/workflows/build.yml) [![GitHub license](https://img.shields.io/github/license/hypersign-protocol/hid-node?color=blue&style=flat-square)](https://github.com/hypersign-protocol/hid-node/blob/main/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hypersign-protocol/hid-node?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/hypersign-protocol/hid-node)](https://goreportcard.com/report/github.com/hypersign-protocol/hid-node) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/hypersign-protocol/hid-node)

<a href="https://discord.gg/CCjJPUuVUz"><img src="https://img.shields.io/discord/308323056592486420?logo=discord" alt="Discord"></a>
<a href="https://twitter.com/intent/follow?screen_name=hypersignchain"> <img src="https://img.shields.io/twitter/follow/hypersignchain?style=social&logo=twitter" alt="follow on Twitter"></a>


The Hypersign Identity Network is a permissionless blockchain network to manage digital identity and access rights. It aims to empower humans to gain control of their data and access on the internet by providing scalable, interoperable and secure [verifiable data registry (VDR)](https://www.w3.org/TR/did-core/#dfn-verifiable-data-registry) to implement use cases on Self Sovereign Identity (SSI) principles. The Hypersign Identity Network is built using [Cosmos-SDK](https://tendermint.com/sdk/) and is fully compatible with [W3C DID specifications](https://www.w3.org/TR/did-core/).

## Features

- Register, Update and Deactivate DID Documents
- Store/Update Credential Schema
- Store/Update status of a Verifiable Credential
- Stake `$HID` tokens
- Submit Governance Proposals
- Transfer `$HID` tokens within and across different Tendermint-based blockchains using IBC
- Deploy CosmWasm Smart Contracts

## Prerequisite

Following are the prerequisites that needs to be installed:

- Golang (Installation Guide: https://go.dev/doc/install) (version: 1.21)
- make
- jq

## Get started

### Local Binary

1. Clone this repository and install the binary:
   ```sh
   git clone https://github.com/hypersign-protocol/hid-node.git
   cd hid-node
   make install
   ```

> The binary `hid-noded` is usually generated in `$HOME/go/bin` directory. Run `hid-noded --help` to explore its functionalities

2. Run the following script to setup a single-node blockchain. Please note that the following script requires `jq` to be installed.
   ```sh
   bash ./scripts/localnet-single-node/setup.sh
   ```

3. Start `hid-noded`:
   ```sh
   hid-noded start
   ```

### Docker

To run a single node `hid-node` docker container, follow the below steps:

1. Pull the image:
   ```sh
   docker pull ghcr.io/hypersign-protocol/hid-node:latest
   ```

2. Run the following:
   ```sh
   docker run --rm -d \
	-p 26657:26657 -p 1317:1317 -p 26656:26656 -p 9090:9090 \
	--name hid-node-container \
	ghcr.io/hypersign-protocol/hid-node start
   ```

## Documentation

| Topic | Reference |
| ----- | ---- |
| Decentralised Identifiers | https://docs.hypersign.id/self-sovereign-identity-ssi/decentralized-identifier-did |
| Credential Schema | https://docs.hypersign.id/self-sovereign-identity-ssi/schema |
| Verifiable Credential Status | https://docs.hypersign.id/self-sovereign-identity-ssi/verifiable-credential-vc/credential-revocation-registry |


Please contact [support@hypermine.in](mailto:support@hypermine.in) for consulting and integration
