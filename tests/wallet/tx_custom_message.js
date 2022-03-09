const { DirectSecp256k1HdWallet, Registry } = require("@cosmjs/proto-signing");
const {
  defaultRegistryTypes,
  SigningStargateClient,
} = require("@cosmjs/stargate");

// Note: The following tx.js being reffered here follows old javascript: vue/src/store/generated/hypersign-protocol/hid-node/hypersignprotocol.hidnode.did/module/types/did/tx.js
// The refactore tx.js file is the `wallet` folder of this repo
const { MsgCreateDID } =  require("./tx.js");

const runfn = async () => {
    const myRegistry = new Registry(defaultRegistryTypes);
    myRegistry.register("/hypersignprotocol.hidnode.ssi.MsgCreateDID", MsgCreateDID); // Replace with your own type URL and Msg class
    const mnemonic = // Replace with your own mnemonic
    "crystal marble excuse boil vendor festival subject grape spatial absorb jaguar keep harbor pass argue fame easy borrow slide exhaust honey clutch attitude slab";

    // Inside an async function...
    // const signer = await DirectSecp256k1HdWallet.fromMnemonic(
    // mnemonic,
    // { prefix: "cosmos" }, // Replace with your own Bech32 address prefix
    // );
    // const client = await SigningStargateClient.connectWithSigner(
    // "http://localhost:26657", // Replace with your own RPC endpoint
    // signer,
    // { registry: myRegistry },
    // );
    
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic);
    const rpcEndpoint = "http://localhost:26657";
    const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, wallet, { registry: myRegistry });
    
    //////////
    const [firstAccount] = await wallet.getAccounts();
    const myAddress = firstAccount.address
    console.log(myAddress)
    const message = {
    typeUrl: "/hypersignprotocol.hidnode.ssi.MsgCreateDID", // Same as above
    value: MsgCreateDID.fromPartial({
        didDocString: "{\"@context\":[\"https://www.w3123.org/ns/did/v1\",\"https://w3id.org/security/v1\",\"https://schema.org\"],\"@type\":\"https://schema.org/Person\",\"id\":\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51\",\"name\":\"Vishwas\",\"publicKey\":[{\"@context\":\"https://w3id.org/security/v2\",\"id\":\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\",\"type\":\"Ed25519VerificationKey2018\",\"publicKeyBase58\":\"5igPDK83gGECDtkKbRNk3TZsgPGEKfkkGXYXLQUfHcd2\"}],\"authentication\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"assertionMethod\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"keyAgreement\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"capabilityInvocation\":[\"did:hs:0f49341a-20ef-43d1-bc93-de30993e6c51#z6MkjAwRoZNV1oifLPb2GzLatZ7sVxY5jZ16xYTTAgSgCqQQ\"],\"created\":\"2021-04-06T14:13:14.018Z\",\"updated\":\"2021-04-06T14:13:14.018Z\"}",
        signatures: [],
        creator: myAddress,
    }),
    };
    const fee = {
    amount: [
        {
        denom: "uatom", // Use the appropriate fee denom for your chain
        amount: "1",
        },
    ],
    gas: "162000",
    };

    // Inside an async function...
    // This method uses the registry you provided
    const response = await client.signAndBroadcast(myAddress, [message], fee);
    console.log(response);
    /** Response Output
        {
        height: 2862,
        transactionHash: '90ED03916FCBEDF03928E984F9B5361C56625395EDA69C7A22FEE86B5FD28720',
        rawLog: '[{"events":[{"type":"message","attributes":[{"key":"action","value":"create_did"}]}]}]',
        data: [ { msgType: '/hypersignprotocol.hidnode.did.MsgCreateDID' } ]
        }
     * **/
}

runfn();