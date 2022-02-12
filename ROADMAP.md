## v0.0.1

**Network Components**

- [ ] Create_DID RPC
- [ ] Update_DID RPC
- [ ] Resolve_DID RPC
- [ ] Create_Schema RPC
- [ ] Query_Schema RPC

**Wallet and Clients**

- [ ] Use Go-CLI to call these RPCs
- [ ] Figure out how a cosmos wallet can be attached with hs-ssi-sdk
- [ ] Upgrade `crypto-ld` package to incorporate `Ed25519Signature2020` instead of `Ed25519Signature2018`
- [ ] Implement JSON rpc with [open-rpc](https://open-rpc.org/) spec 
- [ ] Use JSON rpc in the hs-ssi-sdk (concept of provider need to be implmented) (just like cosmJs use)

**Research**

- [ ] Understand EIP712, EIP1812, EIP2844
- [ ] Read and get complete understanding of IBC whitepaper

## v0.0.2

**Network Components**

- [ ] Work on checking if everything (didDoc, schema etc) is properly align with W3c spec. Do changes if required. 
- [ ] Implment revocation list functionality 
- [ ] Make blockchian explorer with ssi features

**Wallet and Clients**

- [ ] Update the hs-ssi-sdk with v0.0.2
- [ ] Need to find out how we can connect Kepler wallet with hs-ssi-sdk

**Research**

- [ ] _Web3 compatibility  (user should be able to metamask to create DID)_ ?????
- [ ] _Figure out how did RPCs can be called web3?_ ?????


## Future Release

- [ ] Hyperledger aeris compatibility with  Hypersign network
- [ ] Web3 js & Metamask compatibility with Hypersign network 
- [ ] Develop a stateless very light weight client for hyerpsign entwork for developers just like testRPC in ethereum
- [ ] Hypersign SSI playground using webassembly 
- [ ] Develop a thin client (see issue $)


