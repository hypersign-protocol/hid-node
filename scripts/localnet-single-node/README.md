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

- To stop the `hid-node`, press `Ctrl + C`
