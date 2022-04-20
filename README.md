# Hypersign Identity Network

The Hypersign Identity Network is a permissionless blockchain network to manage digital identity and access rights. It aims to empower humans to gain control of their data and access on the internet by providing scalable, interoperable and secure [verifiable data registry (VDR)](https://www.w3.org/TR/did-core/#dfn-verifiable-data-registry) to implement use cases on Self Sovereign Identity (SSI) principles. The Hypersign Identity Network is built using [Cosmos-SDK](https://tendermint.com/sdk/) and is fully compatible with [W3C DID specifications](https://www.w3.org/TR/did-core/).

## Prerequisite

Following are the prerequisites that needs to be installed:

- golang (Installation Guide: https://go.dev/doc/install) (version: 1.17.2)

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

To start a single-node blockchain, refer to the README.md file present in [`/scripts/localnet-single-node/README.md`](/scripts/localnet-single-node/README.md)`

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

### Register DID

```sh
hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
"publicKeyMultibase": "zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
]
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --keyring-backend test --chain-id hidnode
```
Note: While performing a CLI transaction, it is required to pass chain-id as `--chain-id hidnode` , as the default chain id set is `hid-node` which will cause the transaction to fail.

### Update DID

After the DIDDoc is created from running the above command, making changes to it happens through the following CLI command:

```sh
hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"https://some.domain"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
"publicKeyMultibase": "zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
]
}' <version-id> did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --keyring-backend test --chain-id hidnode
```

The second param `<version-id>` should be the version-id of the latest DID Doc.

The `context` field of the DIDDoc is now updated with a new entry: `"https://some.domain"`

### Deactivate DID

Run the following to deactivate the DID Document:

```sh
hid-noded tx ssi deactivate-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52"],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' <version-id> did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --keyring-backend test --chain-id hidnode --yes
```

### Resolve DID

There are two ways to resolve DID:

- CLI
- Blockchain API


**API**:

1. Retrieve a did Document by providing a Did ID:
```sh
curl -X GET "http://localhost:<API-PORT>/hypersign-protocol/hidnode/ssi/did/did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52:" -H  "accept: application/json"
```

2. Retrieve the count and list of did Documents:
```sh
curl -X GET "http://localhost:<API-PORT>/hypersign-protocol/hidnode/ssi/did" -H  "accept: application/json"
```

Note: The above curl command was taken from the Swagger UI of Blockchain API, where the `did` input parameter was entered along with an extra semicolon appended, because gRPC server has issues parsing the regular DID string.

**CLI**:
```sh
hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --chain-id hidnode
```
