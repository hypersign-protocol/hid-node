#!bin/bash

first_test_case () {
    echo "----------------Test case 1: A non-controller DID (User-1) trying to deactivate User-2's DID----------------"
    echo "----------------It should throw an Error----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDocumentMetadata.versionId')
    hid-noded tx ssi deactivate-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from node1 --chain-id hidnode --broadcast-mode block --yes
    echo "----------------x----------------"
}

second_test_case () {
    echo "\n\n----------------Test case 2: Deactivating User-2's DID using its verification key----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDocumentMetadata.versionId')
    hid-noded tx ssi deactivate-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from node1 --chain-id hidnode --broadcast-mode block --yes
    echo "----------------x----------------"
}

third_test_case () {
    echo "\n\n----------------Test case 3: Attempting to update the deactivated DID----------------"
    echo "----------------It should throw an Error----------------"
    VERSION_ID=$(hid-noded query ssi did did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --output json | jq -r '.didDocumentMetadata.versionId')
    hid-noded tx ssi update-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4","did:hs:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' "${VERSION_ID}" did:hs:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from node1 --broadcast-mode block --chain-id hidnode --yes
    echo "----------------x----------------"
}

run_tests() {
    first_test_case
    second_test_case
    third_test_case
}

run_tests