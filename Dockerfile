FROM golang:1.16-alpine

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 jq

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
COPY . .

# Install minimum necessary dependencies and build hid-node
RUN apk add --no-cache $PACKAGES && make build

# Install ca-certificates
RUN apk add --update ca-certificates

# Setup the node
RUN bash ./scripts/docker-node/setup.sh

# Entry for containers, Run the single-node script
ENTRYPOINT [ "bash", "./scripts/docker-node/start.sh" ]

# Expose Ports
EXPOSE 26657 1317 9090 9091 26656
