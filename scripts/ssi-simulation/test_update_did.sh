#!bin/sh

# TODO: Need to use Assertion

echo "----------- Testing for DID Controller -----------"

first_test_case () {
echo "\n\n ------------- TC-1:  Adding User-1's DID to User-2's controller -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:hs:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from node1 --chain-id hidnode --broadcast-mode block --yes

echo "\n\n ------------- TC-1 tested successfully -------------"
}


second_test_case() {
echo "\n\n ------------- TC-2:  Making changes in User-2's DIDDoc with User-1's verification key. (Adding a new element in context field) -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com"
],
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:hs:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from node1 --chain-id hidnode --broadcast-mode block --yes

echo "\n\n ------------- TC-2 tested successfully -------------"
}


third_test_case() {
echo "\n\n ------------- TC-3: User-3 trying to add it's DID in User-2's DIDDoc using its verification key -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com",
"whatelse.com",
"hypersign.id"
],
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:hs:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf","did:hs:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91#key-1 --ver-key JFfT5yPnBbwcDkmry9vdqX9eBKJmfnTT9C1r0LZ5S73BosdZL7AaZ9AYx6Mpvvw/ebaKPyaIiVZ3StijU8RRAA== --from node1 --chain-id hidnode --broadcast-mode block --yes

echo "\n\n ------------- TC-3 tested successfully -------------"
}


fourth_test_case() {
echo "\n\n ------------- TC-4: Adding User-3's DID to User-2's DIDDoc using User-2's verification key -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com"
],
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:hs:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf","did:hs:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from node1 --chain-id hidnode --broadcast-mode block --yes

echo "\n\n ------------- TC-4 tested successfully -------------"
}

run_tests() {
    first_test_case
    second_test_case
    fourth_test_case
    third_test_case
}

run_tests