# SSI Module E2E Tests

Following scenarios are covered for E2E testing:

`TC-1`: A simple SSI flow where three elements of SSI (DID Document, Schema Document and Credential Status Document).<br>
`TC-2`: Controller DID attempts to create Credential Schema and Credential Status Documents on behalf of Parent DID.<br>
`TC-3`: Multiple Controller DID attempt to create Credential Schema and Credential Status Documents on behalf of Parent DID.<br>
`TC-4`: Non-Controller DID attempts to create Credential Schema and Credential Status Documents on behalf of Parent DID (Invalid Case).<br>
`TC-5`: Non-Controller DID attempts to update a DID Document (Invalid Case).<br>
`TC-6`: Controller DID attempts to update the Parent DID Document.<br>
`TC-7`: Parent DID adds multiple DIDs in its controller group, and removes itself. One of the controllers and the Parent DID attemps to change the DID Document.<br>
`TC-8`: Deactivated DID attempts to create Schema and Credential Status Documents.<br>
`TC-9`: `x/ssi` module related transactions, using `secp256k1` keypair.<br>
`TC-10`: Test scenarios for `blockchainAccountId`.<br>
`TC-11`: `x/ssi` module related transactions, using ethereum based `secp256k1` keypair.<br>

## Run Tests

Run the following to run tests
```
# Make sure you are in `./tests/e2e/ssi_tests` directory

python3 run.py
```