const { DirectSecp256k1HdWallet } = require("@cosmjs/proto-signing");
const { SigningStargateClient, StargateClient } = require("@cosmjs/stargate");

const sendCoins = async () => {
  const mnemonic = "crystal marble excuse boil vendor festival subject grape spatial absorb jaguar keep harbor pass argue fame easy borrow slide exhaust honey clutch attitude slab";
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic);
  const [firstAccount] = await wallet.getAccounts();

  const rpcEndpoint = "http://localhost:26657";
  const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, wallet);

  const recipient = "cosmos1zd62yap0fvsy2xdvdjnqjva3qvetwrh0w99a0g";
  const amount = [{
    denom: "uatom",
    amount: "100",
  }];

  const fee = {
    amount: [
        {
        denom: "uatom", 
        amount: "10",
        },
    ],
    gas: "62000",
  };

  console.log({
    addr: firstAccount.address,
    recipient,
    amount: [amount]
  });


  const result = await client.sendTokens(firstAccount.address, recipient, amount, fee,"Have fun with your star coins");
  console.log(result)
};

sendCoins();