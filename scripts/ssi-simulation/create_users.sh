#!/bin/bash

echo "----------- Creating User 1 -----------"

hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
"controller": [],
"alsoKnownAs": ["did:hid:devnet:1f49341a-de30993e6c51"],
"verificationMethod": [
{
"id": "did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf",
"publicKeyMultibase": "zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf"
}
],
"authentication": [
"did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1"
],
"assertionMethod": [
"did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1"
]
}' did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf#key-1 --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from node1 --chain-id hidnode --broadcast-mode block --keyring-backend test --yes

echo "\n----------- User 1 has been created ----------- "

hid-noded query ssi did did:hid:devnet:zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf

echo "--------------- x ----------- "

echo "\n\n ----------- Creating User 2 -----------"

hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"controller": ["did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4",
"publicKeyMultibase": "z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
}
],
"authentication": [
"did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
],
"assertionMethod": [
"did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1"
]
}' did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4#key-1 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from node1 --chain-id hidnode --broadcast-mode block --keyring-backend test --yes

echo "\n ----------- User 2 has been created ----------- "

hid-noded query ssi did did:hid:devnet:z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4

echo "--------------- x ----------- "

echo "\n\n ----------- Creating User 3 -----------"

hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91",
"controller": [],
"alsoKnownAs": [],
"verificationMethod": [
{
"id": "did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91#key-1",
"type": "Ed25519VerificationKey2020",
"controller": "did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91",
"publicKeyMultibase": "zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91"
}
],
"authentication": [
"did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91#key-1"
],
"assertionMethod": [
"did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91#key-1"
]
}' did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91#key-1 --ver-key JFfT5yPnBbwcDkmry9vdqX9eBKJmfnTT9C1r0LZ5S73BosdZL7AaZ9AYx6Mpvvw/ebaKPyaIiVZ3StijU8RRAA== --from node1 --chain-id hidnode --broadcast-mode block --keyring-backend test --yes

echo "\n -----------User 3 has been created ----------- "

hid-noded query ssi did did:hid:devnet:zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91

echo "--------------- x ----------- "
