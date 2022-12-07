# Hypersign Identity Network

[![build-test](https://github.com/hypersign-protocol/vid-node/actions/workflows/build.yml/badge.svg)](https://github.com/hypersign-protocol/vid-node/actions/workflows/build.yml) [![GitHub license](https://img.shields.io/github/license/hypersign-protocol/vid-node?color=blue&style=flat-square)](https://github.com/hypersign-protocol/vid-node/blob/main/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hypersign-protocol/vid-node?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/hypersign-protocol/vid-node)](https://goreportcard.com/report/github.com/hypersign-protocol/vid-node) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/hypersign-protocol/vid-node)

<a href="https://discord.gg/CCjJPUuVUz"><img src="https://img.shields.io/discord/308323056592486420?logo=discord" alt="Discord"></a>
<a href="https://twitter.com/intent/follow?screen_name=hypersignchain"> <img src="https://img.shields.io/twitter/follow/hypersignchain?style=social&logo=twitter" alt="follow on Twitter"></a>


The Hypersign Identity Network is a permissionless blockchain network to manage digital identity and access rights. It aims to empower humans to gain control of their data and access on the internet by providing scalable, interoperable and secure [verifiable data registry (VDR)](https://www.w3.org/TR/did-core/#dfn-verifiable-data-registry) to implement use cases on Self Sovereign Identity (SSI) principles. The Hypersign Identity Network is built using [Cosmos-SDK](https://tendermint.com/sdk/) and is fully compatible with [W3C DID specifications](https://www.w3.org/TR/did-core/).

## Features

- Register, Update and Deactivate DID Documents
- Store/Update Credential Schema
- Store/Update status of a Verifiable Credential
- Stake `$vid` tokens
- Submit Governance Proposals
- Transfer `$vid` tokens within and across different Tendermint-based blockchains

## Prerequisite

Following are the prerequisites that needs to be installed:

- Golang (Installation Guide: https://go.dev/doc/install) (version: 1.18+)
- make

## Get started

### Local

Clone the repository and install the binary:

```sh
git clone https://github.com/hypersign-protocol/vid-node.git
cd vid-node
make install
```

The binary `vid-noded` will be generated in `$GO_PATH/bin` directory. To explore its functionalities, type `vid-noded --help` in a seperate terminal window.

#### Running the Blockchain

To start a single-node blockchain, run the following command to initialize the node:

```sh
sh ./scripts/localnet-single-node/setup.sh
```

> Note: The above script requires `jq` to be installed.

Run the vid-node:

```sh
vid-noded start --home ~/.vid-node
```

### Docker

To run a single node `vid-node` docker container, follow the below steps:

1. Pull the image:
   ```sh
   docker pull ghcr.io/hypersign-protocol/vid-node:latest
   ```

2. Run the following:
   ```sh
   docker run --rm -d \
	-p 26657:26657 -p 1317:1317 -p 26656:26656 -p 9090:9090 \
	--name vid-node-container \
	ghcr.io/hypersign-protocol/vid-node start
   ```

## Documentation

| Topic | Reference |
| ----- | ---- |
| Decentralised Identifiers | https://docs.hypersign.id/self-sovereign-identity-ssi/decentralized-identifier-did |
| Credential Schema | https://docs.hypersign.id/self-sovereign-identity-ssi/schema |
| Verifiable Credential Status | https://docs.hypersign.id/self-sovereign-identity-ssi/verifiable-credential-vc/credential-revocation-registry |

