# LocalNet

## Port Configuration

- RPC: `:26657`
- P2P: `:26656`
- REST Server:`:1317`

## Prerequisite

Install the following:

- golang (ver 1.17+): https://go.dev/doc/install
- jq (`sudo apt-get install jq`)
- tmux (`sudo apt-get install tmux`)

## Running the locanet

- Setup the config files and install `hid-node` binary if it doesn't exists:
  - `sh setup.sh`

- Run the localnet
  - `hid-noded start`

- To display the logs of each node, run the following in a seperate terminal:

Node1: `tmux a -t node1`<br>

## Stop the localnet

```sh
sh stop.sh
```
