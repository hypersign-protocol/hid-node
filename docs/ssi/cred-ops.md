# Verifiable Credential Status

Storing Verifiable Credential on a distributed ledger could lead to privacy violation. However, we can store the status of a Verifiable Credential on-chain, with no private information attached to it. Issuers of a Verifiable Credential have the ability to revoke the credential and provide the reason behind it.

## Register VC Status

For instance, an issuer with id `did:hs:b8da6c12-0833-4c54-af98-55af55c2fd22` has issue a VC, following which they want to register it's status.

CLI Signature is as follow:

```
Usage:
  hid-noded tx ssi register-credential-status [credential-status] [proof]
```

**credential-status Structure**

```json
{
    claim: {
        id: "vc_example1",
            currentStatus: "Live",
            statusReason: "Credential Active"
        },
        issuer: "did:hs:b8da6c12-0833-4c54-af98-55af55c2fd22",
        issued: "2022-04-10T04:07:12Z"
}
```

**proof Structure**

```json
{
    type: "Ed25519VerificationKey2020",
    created: "2022-04-10T04:07:12Z",
    verificationMethod: "did:hs:b8da6c12-0833-4c54-af98-55af55c2fd22#key-1",
    proofValue: "<-- Base64 encoded signature -->",
    proofPurpose: "assertion"
}
```

The field `proofValue` holds the signature that was produced by signing the `credential-status` document. 

### Usage

The following command registers the status of a VC with id `vc_example1`:

```sh
hid-noded tx ssi register-credential-status '{"claim":{"id":"vc_abc1234","currentStatus":"Revoked","statusReason":"It all good"},"issuer":"did:hs:b8da6c12-0833-4c54-af98-55af55c2fd22","issued":"22 May 2022"}' '{"type":"Ed25519VerificationKey2020","created":"2022-04-10T04:07:12Z","verificationMethod":"did:hs:b8da6c12-0833-4c54-af98-55af55c2fd22#key-1","proofValue":"<-- base64 encoded signature -->","proofPurpose":"assertion"}' --from <user-blockchain-address>
```
