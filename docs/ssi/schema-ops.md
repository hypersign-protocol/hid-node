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
 - schema-doc : Schema Document
 - schema-proof : Schema Proof
 - verification-method-id : Verification Method ID
```

## Usage

### Create Schema

Command:

```sh
hid-noded tx ssi create-schema '{"type":"https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json","modelVersion":"v1.0","id":"sch:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf","name":"HS credential template","author":"did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf","authored":"2022-04-10T04:07:12Z","schema":{"schema":"https://json-schema.org/draft-07/schema#","description":"test","type":"object","properties":"{myString:{type:string},myNumner:{type:number},myBool:{type:boolean}}","required":["myString","myNumner","myBool"],"additionalProperties":false}}' '{"type":"Ed25519VerificationKey2020","created":"2022-04-10T04:07:12Z","verificationMethod":"did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1","proofValue":"gLFhwYfObNJEOjNDaeYjprv7FpK0lIhZnFwgOsdRqRHOjQswfm3Hk9EehcYGePrFFwgy4lna73iA5J0BtjfCAw==","proofPurpose":"assertionMethod"}' --from node1 --broadcast-mode block --chain-id hidnode
```

The above command will fail if the User's (`did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf`) DID is not registered on chain

### CLI

1. Query schema for given schema id:

```
http://<REST-URL>/hypersign-protocol/hidnode/ssi/schema/{schemaId}
```

2. Query list of registered schema(s):

```
http://<REST-URL>/hypersign-protocol/hidnode/ssi/schema
```