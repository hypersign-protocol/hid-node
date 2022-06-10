# Schema Operation

## Features

- Transaction Based
    - Create Schema
- Query Based

## CLI Signature

### Create Schema

```
Usage:
  hid-noded tx ssi create-schema [schema] [verification-method-id] [flags]

Params:
 - schema : Schema Document recived from hs-ssi-sdk
 - verification-method-id : Verification Method ID

Flags:
 - --ver-key: Private Key of the signer
```

## Usage

### Create Schema

Command:

```sh
hid-noded tx ssi create-schema '{"type":"https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json","modelVersion":"v1.0","id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51;id=17de181feb67447da4e78259d92d0240;version=1.0","name":"HS credential template","author":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51","authored":"Tue Apr 06 2021 00:09:56 GMT+0530 (India Standard Time)","schema":{"schema":"https://json-schema.org/draft-07/schema#","description":"test","type":"object","properties":"{myString:{type:string},myNumner:{type:number},myBool:{type:boolean}}","required":["myString","myNumner","myBool"],"additionalProperties":false}}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from node1 --keyring-backend test --chain-id hidnode
```

The above command will fail if the User's (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51`) DID is not registered on chain

### CLI

1. Query schema for given schema id:

```
http://<REST-URL>/hypersign-protocol/hidnode/ssi/schema/{schemaId}
```

2. Query list of registered schema(s):

```
http://<REST-URL>/hypersign-protocol/hidnode/ssi/schema
```