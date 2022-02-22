## Tests

### DIDDoc

- Run `create-users.sh` to create three users
- Run `tests/didDoc/test_update_did.sh` to run the tests for DIDdoc Update. Following are the scenarios being tested
  - Adding User-1's DID to User-2's `controller`
  - Making changes in User-2's DIDDoc with User-1's verification key. (Adding a new element in `context` field)
  - User-3 trying to add it's DID in User-2's DIDDoc using its verification key
  - Adding User-3's DID to User-2's DIDDoc using User-2's verification key
- Run `tests/didDoc/test_deactivate_did.sh` to run the tests for DIDdoc Deactivate. Following are the scenarios being tested
  - User-1 (non-controller) trying to deactivate User-2. (It should FAIL)
  - User-2 trying to deactivate to DID using it's own verification key.
  - User-2 attempting to update the DID which is already deactivated. (It should FAIL)