FROM golang:1.19

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc python3 jq

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
COPY . .

# Install minimum necessary dependencies and build hid-node
RUN apt-get update
RUN apt-get install ${PACKAGES} -y
RUN make install 

# Setup the node
RUN bash ./scripts/docker-node/setup.sh

# Expose Ports
EXPOSE 26657 1317 9090 9091 26656

# Entry for containers
ENTRYPOINT [ "hid-noded" ]