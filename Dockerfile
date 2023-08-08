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

# Expose Ports
EXPOSE 26657 1317 9090 9091 26656

# Provide permission to the script
RUN chmod +x ./scripts/docker-node/entrypoint.sh

# Entry for containers
ENTRYPOINT [ "/bin/sh", "-c", "echo N | ./scripts/docker-node/entrypoint.sh" ]

