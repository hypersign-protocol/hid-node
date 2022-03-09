# SSI Module Walkthrough

## Features

- Transaction Based:
  - Registering a DID Document
  - Updating a DID Document
  - Deactivating a DID Document
- Querying Based:
  - Resolve a DID Document based on an input DID Id
  - Get the count and list of DID Documents registered on chain
  
## CLI Signature

### Register DID

```
Usage:
  hid-noded tx ssi create-did [did-doc-string] [verification-method-id] [flags]

Params:
 - did-doc-string : Did Doc String received from hs-ssi-sdk
 - verification-method-id : Verification Method ID

Flags:
 - ver-key : Private Key of the Signer
```

### Update DID

```
Usage:
  hid-noded tx ssi update-did [did-doc-string] [version-id] [verification-method-id] [flags]

Params:
 - did-doc-string : Did Doc String received from hs-ssi-sdk
 - version-id : Version ID of the DID Document to be updated. It is expected that version ID should of the latest DID Document
 - verification-method-id : Verification Method ID

Flags:
 - --ver-key : Private Key of the Signer
```

### Deactivate DID

```
Usage:
  hid-noded tx ssi deactivate-did [did-doc-string] [version-id] [verification-method-id] [flags]

Params:
 - did-doc-string : Did Doc String received from hs-ssi-sdk
 - version-id : Version ID of the DID Document to be updated. It is expected that version ID should of the latest DID Document
 - verification-method-id : Verification Method ID

Flags:
 - --ver-key : Private Key of the Signer
```

## Usage

The usage of CLI is explained through scenarios:

### Register DID

Registering a DID Document in `hid-node`. User 2 is registering a DID Document with id: `did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52`

```sh
hid-noded tx ssi create-did '{
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
"service": [{
"id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#vcs",
"type": "LinkedDomains",
"serviceEndpoint": "https://example.com/vc"
},
{
"id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#file",
"type": "LinkedDomains",
"serviceEndpoint": "https://example.in/somefile"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key <private-key> --from alice --keyring-backend test --chain-id hidnode --yes
```

### Update DID

User 2 (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52`) is trying to update it’s DID by adding User 1’s ID (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51`) to the controller group. It is assumed that User 1’s ID (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51`) is already registered on blockchain.

```sh
hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51"],
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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key <private-key> --keyring-backend test --from alice --chain-id hidnode --yes
```

Here, the `${VERSION_ID}` should have the version id of the latest DID of User 2 (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52`)

### Deactivate DID

User 2 (`did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52`) is trying to deactivate it’s DID

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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key <private-key> --keyring-backend test --from alice --chain-id hidnode --yes
```

### Resolving DID

1) Get the list of Registered DID Documents

URL: `http://localhost:1318/hypersign-protocol/hidnode/ssi/did`

Output:

```json
{
   "totalDidCount":"2",
   "didDocList":[
      {
         "_at_context":"",
         "didDocument":{
            "context":[
               "https://www.w3.org/ns/did/v1",
               "https://w3id.org/security/v1",
               "https://schema.org"
            ],
            "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
            "controller":[
               
            ],
            "alsoKnownAs":[
               "did:hs:1f49341a-de30993e6c51"
            ],
            "verificationMethod":[
               {
                  "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
                  "type":"Ed25519VerificationKey2020",
                  "controller":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
                  "publicKeyMultibase":"zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
               }
            ],
            "authentication":[
               "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
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
               
            ]
         },
         "didDocumentMetadata":{
            "created":"2022-02-25T09:20:15Z",
            "updated":"2022-02-25T09:20:15Z",
            "deactivated":false,
            "versionId":"GkAO5TuRaFWnMD3IgoKaaBMKEIByYWIi9h/W9LvLk+Q="
         },
         "didResolutionMetadata":{
            "retrieved":"2022-02-25T09:20:19Z",
            "error":""
         }
      },
      {
         "_at_context":"",
         "didDocument":{
            "context":[
               "https://www.w3.org/ns/did/v1",
               "https://w3id.org/security/v1",
               "https://schema.org"
            ],
            "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
            "controller":[
               "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52"
            ],
            "alsoKnownAs":[
               "did:hs:1f49341a-de30993e6c52"
            ],
            "verificationMethod":[
               {
                  "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
                  "type":"Ed25519VerificationKey2020",
                  "controller":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
                  "publicKeyMultibase":"z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
               }
            ],
            "authentication":[
               "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
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
                  "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#vcs",
                  "type":"LinkedDomains",
                  "serviceEndpoint":"https://example.com/vc"
               },
               {
                  "id":"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#file",
                  "type":"LinkedDomains",
                  "serviceEndpoint":"https://example.in/somefile"
               }
            ]
         },
         "didDocumentMetadata":{
            "created":"2022-02-25T09:20:11Z",
            "updated":"2022-02-25T09:20:11Z",
            "deactivated":false,
            "versionId":"ClUei1OW9mDtFQuFdhgmfzPZT1gWa7hGwfRI9DP2mMs="
         },
         "didResolutionMetadata":{
            "retrieved":"2022-02-25T09:20:19Z",
            "error":""
         }
      }
   ]
}
```

2) Get the list of Registered DID Documents with pagination limit

URL: `http://localhost:1318/hypersign-protocol/hidnode/ssi/did?pagination.limit=1`

Output:

```json
{
  "totalDidCount": "2",
  "didDocList": [
    {
      "_at_context": "",
      "didDocument": {
        "context": [
          "https://www.w3.org/ns/did/v1",
          "https://w3id.org/security/v1",
          "https://schema.org"
        ],
        "id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
        "controller": [],
        "alsoKnownAs": [
          "did:hs:1f49341a-de30993e6c51"
        ],
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
        ],
        "assertionMethod": [],
        "keyAgreement": [],
        "capabilityInvocation": [],
        "capabilityDelegation": [],
        "service": []
      },
      "didDocumentMetadata": {
        "created": "2022-02-25T15:18:59Z",
        "updated": "2022-02-25T15:18:59Z",
        "deactivated": false,
        "versionId": "OwpjbfvZn5mBdf1gJWrpYFKrI2yLCQAjVhgHCqq6WOo="
      },
      "didResolutionMetadata": {
        "retrieved": "2022-02-25T15:19:05Z",
        "error": ""
      }
    }
  ]
}
```

3) Query the DID Document for a given DID Id

URL: `http://localhost:1318/hypersign-protocol/hidnode/ssi/did/did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52:`

<br>

```Note the colon(:) at the end of URL. It has been appended because of limitations of gRPC Server in parsing the DID Id. Workaround for this is being upon```

<br>

Output: 

```json
{
  "_at_context": "",
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1",
      "https://w3id.org/security/v1",
      "https://schema.org"
    ],
    "id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
    "controller": [
      "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52"
    ],
    "alsoKnownAs": [
      "did:hs:1f49341a-de30993e6c52"
    ],
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
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": [],
    "service": [
      {
        "id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#vcs",
        "type": "LinkedDomains",
        "serviceEndpoint": "https://example.com/vc"
      },
      {
        "id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#file",
        "type": "LinkedDomains",
        "serviceEndpoint": "https://example.in/somefile"
      }
    ]
  },
  "didDocumentMetadata": {
    "created": "2022-02-25T09:20:11Z",
    "updated": "2022-02-25T09:20:11Z",
    "deactivated": false,
    "versionId": "ClUei1OW9mDtFQuFdhgmfzPZT1gWa7hGwfRI9DP2mMs="
  },
  "didResolutionMetadata": {
    "retrieved": "2022-02-25T09:24:43Z",
    "error": ""
  }
}
```