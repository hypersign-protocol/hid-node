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