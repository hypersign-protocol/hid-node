## Tests

### DIDDoc

- Run `create-users.sh` to create three users
- Run `tests/didDoc/test_did.sh` to run the tests for DIDdoc. Following are the scenarios being tested
  - Adding User-1's DID to User-2's `controller`
  - Making changes in User-2's DIDDoc with User-1's verification key. (Adding a new element in `context` field)
  - User-3 trying to add it's DID in User-2's DIDDoc using its verification key
  - Adding User-3's DID to User-2's DIDDoc using User-2's verification key