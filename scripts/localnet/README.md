# LocalNet

The script `start-3node.sh` runs a 3-Node local hid-node network

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

### step1  : Install Golang (ver 1.17.2)



### step 2: Setup `hid-noded` binary 

```bash
wget https://github.com/hypersign-protocol/hid-node/releases/download/latest/hid-node_latest_linux_amd64.tar.gz
tar -xvzf hid-node_latest_linux_amd64.tar.gz
mv hid-noded ~/go/bin
```
### step 3: Install `tmux`


## Running the locanet

- Provide permissions to `start-3node.sh` to execute:
  ```sh
  sudo chmod +x path/to/start-3node.sh
  ```
- Execute the script:
  ```sh
  sh path/to/start-3node.sh
  ```

## Stop  the localnet

```sh
sh path/to/stop-3node.sh
```

To display the logs of each node, run the following:

Node1: `tmux a -t node1`<br>
Node2: `tmux a -t node2`<br>
Node3: `tmux a -t node2`<br>