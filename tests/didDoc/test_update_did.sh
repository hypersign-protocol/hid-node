#!bin/sh

# TODO: Need to use Assertion

echo "----------- Testing for DID Controller -----------"

first_test_case () {
echo "\n\n ------------- TC-1:  Adding User-1's DID to User-2's controller -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')

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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes

echo "\n\n ------------- TC-1 tested successfully -------------"
}


second_test_case() {
echo "\n\n ------------- TC-2:  Making changes in User-2's DIDDoc with User-1's verification key. (Adding a new element in context field) -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com"
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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --chain-id hidnode --yes

echo "\n\n ------------- TC-2 tested successfully -------------"
}


third_test_case() {
echo "\n\n ------------- TC-3: User-3 trying to add it's DID in User-2's DIDDoc using its verification key -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com",
"whatelse.com",
"hypersign.id"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53"],
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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53#zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91 --ver-key JFfT5yPnBbwcDkmry9vdqX9eBKJmfnTT9C1r0LZ5S73BosdZL7AaZ9AYx6Mpvvw/ebaKPyaIiVZ3StijU8RRAA== --from alice --chain-id hidnode --yes

echo "\n\n ------------- TC-3 tested successfully -------------"
}


fourth_test_case() {
echo "\n\n ------------- TC-4: Adding User-3's DID to User-2's DIDDoc using User-2's verification key -------------"

VERSION_ID=$(hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52 --output json | jq -r '.didDoc.metadata.versionId')

hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org",
"something.com"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52",
"controller": ["did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51","did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53"],
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
}' "${VERSION_ID}" did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes

echo "\n\n ------------- TC-4 tested successfully -------------"
}

run_tests() {
    first_test_case
    second_test_case
    fourth_test_case
    third_test_case
}

run_tests