VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git rev-parse --short HEAD)

GOBIN = $(shell go env GOPATH)/bin

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

export GO111MODULE=on

###############################################################################
###                                  Build                                  ###
###############################################################################
.PHONY: build

build: go.sum
		go build -mod=readonly $(BUILD_FLAGS) -o ./build/hid-noded ./cmd/hid-noded		

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-build:
	sh ./scripts/protocgen.sh