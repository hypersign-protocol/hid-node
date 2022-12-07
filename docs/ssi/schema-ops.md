# Schema Operations Walkthrough

> In Progress

## Syntax of Schema ID 

The syntax of Schema ID is as follows:

```
sch:vid:<chain-namespace>:<method-specific-id>:<version-number>
```

- `sch:vid` - Schema Method, where `sch` is the document identifier and `vid` is the method name
- `<chain-namespace>` - *(Optional)* Name of the blockchain where the schema document is registered. It is omitted for the document registered on mainnet chain
- `<method-specific-id>` - Multibase-encoded unique identifier of length 45
- `<version-number>` - Model version of schema. For instance, `1.0`, `1.1` and `2.1`

## Schema Operations

- **Transaction Based**
  - Register/Update a Schema Document
- **Query Based**
  - Query a Schema Document
  - Query Registered Schema Documents

## Usage

### Register/Update Schema

Both registration and update of Schema happens through the RPC `CreateSchema`

**CLI Signature**

```
Usage:
  vid-noded tx ssi create-schema [schema-doc] [schema-proof] [flags]

Params:
 - schema-doc : Schema Document
 - schema-proof : Schema Proof
```

**Example**

```sh
vid-noded tx ssi create-schema '{"type":"https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json","modelVersion":"v1.0","id":"sch:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf:1.0","name":"HS credential template","author":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf","authored":"2022-04-10T04:07:12Z","schema":{"schema":"https://json-schema.org/draft-07/schema#","description":"test","type":"object","properties":"{myString:{type:string},myNumner:{type:number},myBool:{type:boolean}}","required":["myString","myNumner","myBool"],"additionalProperties":false}}' '{"type":"Ed25519VerificationKey2020","created":"2022-04-10T04:07:12Z","verificationMethod":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1","proofValue":"gLFhwYfObNJEOjNDaeYjprv7FpK0lIhZnFwgOsdRqRHOjQswfm3Hk9EehcYGePrFFwgy4lna73iA5J0BtjfCAw==","proofPurpose":"assertionMethod"}' --from <key-name-or-address> --chain-id <Chain ID> --yes
```

### Register/Update Schema

**CLI Signature**

```
Usage:
  vid-noded query ssi schema [schema-id]

Params:
 - schema-id : Schema ID
```

**Example**

```
vid-noded query ssi schema sch:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf:1.0
```

**REST**

1. Query a schema document for given schema id:

```
http://<REST-URL>/hypersign-protocol/vidnode/ssi/schema/{schemaId}
```

2. Query a list of registered schema documents:

```
http://<REST-URL>/hypersign-protocol/vidnode/ssi/schema
```