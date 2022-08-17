# Hypersign Identity Network

The Hypersign Identity Network is a permissionless blockchain network to manage digital identity and access rights. It aims to empower humans to gain control of their data and access on the internet by providing scalable, interoperable and secure [verifiable data registry (VDR)](https://www.w3.org/TR/did-core/#dfn-verifiable-data-registry) to implement use cases on Self Sovereign Identity (SSI) principles. The Hypersign Identity Network is built using [Cosmos-SDK](https://tendermint.com/sdk/) and is fully compatible with [W3C DID specifications](https://www.w3.org/TR/did-core/).

## Prerequisite

Following are the prerequisites that needs to be installed:

- golang (Installation Guide: https://go.dev/doc/install) (version: 1.18.5)

## Get started

### Local:

Clone the hid-node repository and build the binary:

```sh
git clone https://github.com/hypersign-protocol/hid-node.git
cd hid-node
make build
```

The binary `hid-noded` will be generated in `$GO_PATH/bin` directory. To explore its functionalities, type `hid-noded --help` im a seperate terminal window.

#### Running the Blockchain

To start a single-node blockchain, run the following command to initialize the node:

```sh
sh ./scripts/localnet-single-node/setup.sh
```
Run the hid-node:

```sh
hid-noded start --home ~/.hid-node
```

### Docker:

To run a single node `hid-node` docker container, run the following:

1. Pull the image:
   ```sh
   docker pull hypersignprotocol/hid-node
   ```

2. Open a separate terminal window. Run the node:
   ```sh
   docker run -it hypersignprotocol/hid-node start
   ```

## Operations

Hands-on CLI operations are present in the below docs:

- [Decentralised Identifier (DID)](docs/ssi/did-ops.md)
- [Credential Schema](docs/ssi/schema-ops.md)
- [Verifiable Credential Status](docs/ssi/cred-ops.md)