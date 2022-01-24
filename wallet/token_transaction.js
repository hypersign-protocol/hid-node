const { DirectSecp256k1HdWallet } = require("@cosmjs/proto-signing");
const { assertIsBroadcastTxSuccess, SigningStargateClient, StargateClient } = require("@cosmjs/stargate");

const sendCoins = async () => {
  const mnemonic = "cactus rail narrow minute human cannon mother subway decide endless vicious boss sister claw lawn swap orbit have end van rose alcohol wire diary";
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic);
  const [firstAccount] = await wallet.getAccounts();

  const rpcEndpoint = "http://localhost:26657";
  const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, wallet);

  const recipient = "cosmos19ttzz9980jefqymdwv4v0h80j0qu4ehz74j3u2";
  const amount = {
    denom: "uatom",
    amount: "20",
  };

  console.log({
    addr: firstAccount.address,
    recipient,
    amount: [amount]
  });


  const result = await client.sendTokens(firstAccount.address, recipient, [amount],"Have fun with your star coins");
  console.log(result)
  assertIsBroadcastTxSuccess(result);
};

sendCoins();