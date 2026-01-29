FROM golang:1.21

# Set up dependencies
# ENV PACKAGES curl make git libc-dev bash gcc python3 jq

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
COPY . .

# Install minimum necessary dependencies and build hid-node
# RUN apt-get update
# RUN apt-get install ${PACKAGES} -y
RUN make install 

# Setup the node
# RUN bash ./scripts/docker-node/setup.sh
# Expose Ports
EXPOSE 26657 1317 9090 9091 26656 1234 26660
# Entry for containers

COPY ./genesis.json ./genesis.json


CMD   mv ./genesis.json /root/.hid-node/config/genesis.json ; sed -i -E '119s/enable = false/enable = true/' /root/.hid-node/config/app.toml ; sed -i -E '122s/swagger = false/swagger = true/' /root/.hid-node/config/app.toml; sed -i -E '140s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' /root/.hid-node/config/app.toml ;sed -i -E 's|tcp://localhost:1317|tcp://0.0.0.0:1317|g' /root/.hid-node/config/app.toml; sed -i -E 's|localhost:9090|0.0.0.0:9090|g' /root/.hid-node/config/app.toml; sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /root/.hid-node/config/config.toml; sed -i 's/priv_validator_laddr = ""/priv_validator_laddr = "tcp:\/\/0.0.0.0:1234"/' /root/.hid-node/config/config.toml ; sed -i -E 's|prometheus = false|prometheus = true|g' /root/.hid-node/config/config.toml ;sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /root/.hid-node/config/config.toml; sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /root/.hid-node/config/config.toml;  hid-noded start;

