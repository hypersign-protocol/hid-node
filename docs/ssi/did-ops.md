# DID Operation Walkthrough

Decentralised Identifiers are cryptographically-verifiable identifiers which are stored on a decentralised ledger, which enables users to own and manage their ID.

## Syntax of `did:vid` method

The `did:vid` method are as follows:

```
did                = "did:" method-name ":" [chain-namespace] ":" method-specific-id
method-name        = "vid"
chain-namespace    = ALPHA / DIGIT
method-specific-id = 45 * id-char
id-char            = ALPHA / DIGIT
```

**Description of ID segments**

- `did` - Document Identifier of DID Document
- `vid` - Method name
- `<chain-namespace>` - *(Optional)* Name of the blockchain where the VC status is registered. It is omitted for the document registered on mainnet chain
- `<method-specific-id>` - Multibase-encoded unique identifier of length 45

## Supported DID Method Operations

The `did:vid` method supports the following operations:

- **Transaction Based**:
  - Register a DID document
  - Update a DID document
  - Deactivate a DID document
- **Query Based**:
  - Query a DID Document
  - Query registered DID Documents

## Usage

### Register DID

**CLI Signature**

```
Usage:
  vid-noded tx ssi create-did [did-doc-string] [verification-method-id] [flags]

Params:
 - did-doc-string : Did Document String
 - verification-method-id : Id of verification Method Key

Flags:
 - ver-key : Private Key of the Signer
```

**Example**

```sh
vid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",`
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"],
"alsoKnownAs": ["did:vid:devnet:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"service": [{
"id":"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#vcs",
"type": "LinkedDomains",
"serviceEndpoint": "http://example.onion"
},
{
"id":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#file",
"type": "LinkedDomains",
"serviceEndpoint": "http://service.onion"
}
],
"authentication": [
"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
],
"assertionMethod": [
"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key <base64-encoded-private-key> --from <key-name-or-address> --chain-id <Chain ID> --yes
```

### Query DID

**CLI Signature**

```
Usage:
  vid-noded q ssi did [didDoc-id] [flags]

Params:
 - didDoc-id : Did Document Id
```

**Example**

```
vid-noded q ssi did did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4
```

**REST**

1) Get the list of registered DID Documents

URL: `http://<REST-URL>/hypersign-protocol/vidnode/ssi/did`

Output:

```json
{
   "totalDidCount":"2",
   "didDocList":[
      {
         "didDocument":{
            "context":[
               "https://www.w3.org/ns/did/v1"
            ],
            "id":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
            "controller":[
               "did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
            ],
            "alsoKnownAs":[
               "did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
            ],
            "verificationMethod":[
               {
                  "id":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1",
                  "type":"Ed25519VerificationKey2020",
                  "controller":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
                  "publicKeyMultibase":"zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
               }
            ],
            "authentication":[
               "did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1"
            ],
            "assertionMethod":[
               "did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1"
            ],
            "keyAgreement":[
               
            ],
            "capabilityInvocation":[
               
            ],
            "capabilityDelegation":[
               
            ],
            "service":[
               
            ]
         },
         "didDocumentMetadata":{
            "created":"2022-02-25T09:20:15Z",
            "updated":"2022-02-25T09:20:15Z",
            "deactivated":false,
            "versionId":"BC2B9B37AF99BC7BC5599124B2F2168066B1B2DDE8F8C06FD720B3DDF2803285"
         }
      },
      {
         "didDocument":{
            "context":[
               "https://www.w3.org/ns/did/v1"
            ],
            "id":"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
            "controller":[
               "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
            ],
            "alsoKnownAs":[
               "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
            ],
            "verificationMethod":[
               {
                  "id":"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
                  "type":"Ed25519VerificationKey2020",
                  "controller":"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
                  "publicKeyMultibase":"z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
               }
            ],
            "authentication":[
               "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
            ],
            "assertionMethod":[
               
            ],
            "keyAgreement":[
               
            ],
            "capabilityInvocation":[
               
            ],
            "capabilityDelegation":[
               
            ],
            "service":[
               {
                  "id":"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#vcs",
                  "type":"LinkedDomains",
                  "serviceEndpoint":"https://example.com/vc"
               },
               {
                  "id":"did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#file",
                  "type":"LinkedDomains",
                  "serviceEndpoint":"https://example.in/somefile"
               }
            ]
         },
         "didDocumentMetadata":{
            "created":"2022-02-25T09:20:11Z",
            "updated":"2022-02-25T09:20:11Z",
            "deactivated":false,
            "versionId":"72093DE93BE67801C160C7CA12DC844D63D9FB3922F3ECC94BBB27844DE568D1"
         }
      }
   ]
}
```

2) Query the DID Document for an input DID id

URL: `http://<REST-URL>/hypersign-protocol/vidnode/ssi/did/did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4:`

```Note the colon(:) at the end of URL. It has been appended because of limitations of gRPC Server in parsing the DID Id. Workaround for this is being upon```

Output: 

```json
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
    "controller": [
      "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
    ],
    "alsoKnownAs": [
      "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
    ],
    "verificationMethod": [
      {
        "id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
        "type": "Ed25519VerificationKey2020",
        "controller": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
        "publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
      }
    ],
    "authentication": [
      "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": [],
    "service": [
      {
        "id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#vcs",
        "type": "LinkedDomains",
        "serviceEndpoint": "https://example.com/vc"
      },
      {
        "id": "did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#file",
        "type": "LinkedDomains",
        "serviceEndpoint": "https://example.in/somefile"
      }
    ]
  },
  "didDocumentMetadata": {
    "created": "2022-02-25T09:20:11Z",
    "updated": "2022-02-25T09:20:11Z",
    "deactivated": false,
    "versionId": "72093DE93BE67801C160C7CA12DC844D63D9FB3922F3ECC94BBB27844DE568D1"
  }
}
```

### Update DID

**CLI Signature**

```
Usage:
  vid-noded tx ssi update-did [did-doc-string] [version-id] [verification-method-id] [flags]

Params:
 - did-doc-string : Did Document string
 - version-id : Version ID of the DID Document to be updated. It is expected that version Id should match latest DID Document's version Id
 - verification-method-id : Id of verification Method Key

Flags:
 - --ver-key : Private Key of the Signer
```

**Example**

```sh
vid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:vid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"],
"alsoKnownAs": ["did:vid:devnet:1f49341a-de30993e6c52"],
"verificationMethod": [
{
"id": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' <version-id> did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key <private-key> --from <key-name-or-address> --chain-id <Chain ID> --yes
```

### Deactivate DID

**CLI Signature**

```
Usage:
  vid-noded tx ssi deactivate-did [did-id] [version-id] [verification-method-id] [flags]

Params:
 - did-id : DID Document ID
 - version-id : Version ID of the DID Document to be deactivated. It is expected that version Id should match latest DID Document's version Id
 - verification-method-id : Id of verification Method Key

Flags:
 - --ver-key : Private Key of the Signer
```

**Example**

```sh
vid-noded tx ssi deactivate-did 'did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4' <version-id> did:vid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key <private-key> --from <key-name-or-address> --chain-id <Chain Id> --yes
```
