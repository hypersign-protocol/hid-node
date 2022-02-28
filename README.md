# Hypersign Identity Network
The  Hypersign Identity Network is a permissionless blockchain network to manage digital identity and access rights. It aims to empower humans to gain control of their data and access on the internet by providing scalable, interoperable and secure verifiable data registry (VDR) to implement use cases on Self Sovereign Identity (SSI) principles. The Hypersign Identity Network is built using Cosmos-SDK and is fully compatible with W3C DID specification.

## Prerequisite

Following are the prerequisites that needs to be installed:

- golang (Installation Guide: https://go.dev/doc/install) (version: 1.17.2)
- starport (Installation Guide: https://docs.starport.network/guide/install.html)

## Get started

Clone the hid-node repository:

```
$ git clone https://github.com/hypersign-protocol/hid-node.git
$ cd hid-node
```

Run the following command to build the binary file and start the blockchain: 
```
starport chain serve
```

You now have a blockchain up and running!

The binary `hid-noded` will be generated in `$GO_PATH/bin` directory. To explore its functionalities, type `hid-noded --help` im a seperate terminal window.

To stop the blockchain, navigate to the terminal window where the blockchain is running, and hit `Ctrl+C`.

## Register DID

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
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --chain-id hidnode
```
Note: While performing a CLI transaction, it is required to pass chain-id as `--chain-id hidnode` , as the default chain id set is `hid-node` which will cause the transaction to fail.

## Operations

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
}' <version-id> did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --chain-id hidnode
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
}' <version-id> did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes
```

## Resolve DID

There are two ways to resolve DID:

- CLI
- Blockchain API


**API**:

1. Retrieve a did Document by providing a Did ID:
```sh
curl -X GET "http://localhost:1318/hypersign-protocol/hidnode/ssi/did/did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52:" -H  "accept: application/json"
```

2. Retrieve the count and list of did Documents:
```sh
curl -X GET "http://localhost:1318/hypersign-protocol/hidnode/ssi/did" -H  "accept: application/json"
```

Note: The above curl command was taken from the Swagger UI of Blockchain API, where the `did` input parameter was entered along with an extra semicolon appended, because gRPC server has issues parsing the regular DID string.

**CLI**:
```sh
hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52
```