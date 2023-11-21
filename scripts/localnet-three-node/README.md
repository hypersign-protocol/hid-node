# LocalNet

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

### step 1: Install Golang (ver 1.17.2)

### step 2: Setup `hid-noded` binary 

```bash
wget https://github.com/hypersign-protocol/hid-node/releases/download/latest/hid-node_latest_linux_amd64.tar.gz
tar -xvzf hid-node_latest_linux_amd64.tar.gz
mv hid-noded ~/go/bin
```
### step 3: Install `tmux` and `jq`


## Running the locanet

- Execute the script to execute the complete workflow:
  ```sh
  sh all.sh
  ```

## Stop  the localnet

```sh
sh path/to/stop.sh
```

To display the logs of each node, run the following:

Node1: `tmux a -t node1`<br>
Node2: `tmux a -t node2`<br>
Node3: `tmux a -t node2`<br>