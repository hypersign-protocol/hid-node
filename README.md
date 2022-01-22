# hid-node
**hid-node** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

## Prerequisite

Following are the prerequisites that needs to be installed:

- golang (Installation Guide: https://go.dev/doc/install)
- starport (Installation Guide: https://docs.starport.network/guide/install.html)

## Get started

Clone the hid-node repository:

```
$ git clone https://github.com/hypersign-protocol/hid-node.git
$ cd hid-node
```

Run the following command to build the binary file and start the `hid-node` blockchain: 
```
starport chain serve
```

You now have a blockchain up and running!

The binary `hid-noded` will be generated in `$GO_PATH/bin` directory. To explore its functionalities, type `hid-noded --help` im a seperate terminal window.

To stop the blockchain, navigate to the terminal window where the blockchain is running, and hit `Ctrl+C`.

## Module Creation

Once we have scaffolded the chain using `starport`, a default module is always created if we scaffolded the chain without the `--no-module` flag. In our case, it will be `x/hidnode`. We can delete this module and its dependent files and folders, since it's not necessary.

Creating the module `did` is as follows:

```
$ starport scaffold module did
```

Now, to scaffold any structures such as `messages`, `types`, `list` etc., following needs to run:

```
$ starport scaffold message createDID did didDocString createdAt --module did
```
Notice the `--module` flag. This is required to specify the module for which we are scaffolding the structure


## Register DID

```
hid-noded tx did create-did "did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51" "{\"@context\":[\"https://www.w3.org/ns/did/v1\",\"https://w3id.org/security/v1\",\"https://schema.org\"],\"@type\":\"https://schema.org/Person\",\"id\":\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51\",\"name\":\"Vishwas\",\"publicKey\":[{\"@context\":\"https://w3id.org/security/v2\",\"id\":\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\",\"type\":\"Ed25519VerificationKey2018\",\"publicKeyBase58\":\"5igPDK83gGECDtkKbRNk3TZsgPGEKfkkGXYXLQUfHcd2\"}],\"authentication\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"assertionMethod\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"keyAgreement\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"capabilityInvocation\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"created\":\"2021-04-06T14:13:14.018Z\",\"updated\":\"2021-04-06T14:13:14.018Z\"}" "2022-01-21T13:03:22.023Z" --from alice --chain-id hidnode
```
Note: While performing a CLI transaction, it is required to pass chain-id as `--chain-id hidnode` , as the default chain id set is `hid-node` which will cause the transaction to fail.

## Resolve DID

```
curl -X GET "http://localhost:1318/hypersign-protocol/hidnode/did/did%3Ahs%3A0f49341a-20ef-43d1-bc93-de30993e6c51%3A" -H  "accept: application/json"
```
Note: The above curl command was taken from the Swagger UI of Blockchain API, where the `did` input parameter was entered along with an extra semicolon appended, because gRPC server has issues parsing the regular DID string.