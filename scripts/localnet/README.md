# LocalNet

The script `localnet-three-node.sh` runs a 3-Node local hid-node network

## Port Configuration

- Node 1
  - RPC: `:26657`
  - P2P: `:26656`
  - REST Server:`:1317`

- Node 2
  - RPC: `:36657`
  - P2P: `:36656`
  - REST Server:`:2317`

- Node 3
  - RPC: `:46657`
  - P2P: `:46656`
  - REST Server:`:3317`

## Prerequisite

- Golang (ver 1.17.2)
- `hid-node` binary already set-up

## Running the locanet

- Provide permissions to `localnet-three-node.sh` to execute:
  ```sh
  sudo chmod +x path/to/localnet-three-node.sh
  ```
- Execute the script:
  ```sh
  sh path/to/localnet-three-node.sh
  ```

To display the logs of each node, run the following:

Node1: `tmux a -t node1`<br>
Node2: `tmux a -t node2`<br>
Node3: `tmux a -t node2`<br>