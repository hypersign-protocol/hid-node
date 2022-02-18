#!bin/sh

echo "----------- Creating User 1 -----------"

hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51",
"controller": [],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c51"],
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
]
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#zEYJrMxWigf9boyeJMTRN4Ern8DJMoCXaLK77pzQmxVjf --ver-key oVtY1xceDZQjkfwlbCEC2vgeADcxpgd27vtYasBhcM/JLR6PnPoD9jvjSJrMsMJwS7faPy5OlFCdj/kgLVZMEg== --from alice --chain-id hidnode --yes

echo "\n----------- User 1 has been created ----------- "

hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51

echo "--------------- x ----------- "

echo "\n\n ----------- Creating User 2 -----------"

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
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4"
]
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52#z8BXg2zjwBRTrjPs7uCnkFBKrL9bPD14HxEJMENxm3CJ4 --ver-key bZBUkLGChnJujYHUZ4L8PECoN2Odv6adWGXc1qVWCRVqtEx0o/FmtFZnd5pT3laR518P58TRUGY5q5KSrToSmQ== --from alice --chain-id hidnode --yes

echo "\n ----------- User 2 has been created ----------- "

hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c52

echo "--------------- x ----------- "

echo "\n\n ----------- Creating User 3 -----------"

hid-noded tx ssi create-did '{
"context": [
"https://www.w3.org/ns/did/v1",
"https://w3id.org/security/v1",
"https://schema.org"
],
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53",
"controller": [],
"alsoKnownAs": ["did:hs:1f49341a-de30993e6c53"],
"verificationMethod": [
{
"id": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53#zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91",
"type": "Ed25519VerificationKey2020",
"controller": "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53",
"publicKeyMultibase": "zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91"
}
],
"authentication": [
"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53#zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91"
]
}' did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53#zE2seaoaAwBzfLaTd7oqEf5GJZdEiQgo64ayJgMstRZ91 --ver-key JFfT5yPnBbwcDkmry9vdqX9eBKJmfnTT9C1r0LZ5S73BosdZL7AaZ9AYx6Mpvvw/ebaKPyaIiVZ3StijU8RRAA== --from alice --chain-id hidnode --yes

echo "\n -----------User 3 has been created ----------- "

hid-noded query ssi did did:hs:0f49341a-20ef-43d1-bc93-de30993e6c53

echo "--------------- x ----------- "