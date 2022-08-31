#!/usr/bin/make -f

VERSION := $(shell git describe --tags --abbrev=0)
COMMIT := $(shell git rev-parse --short HEAD)

GOBIN = $(shell go env GOPATH)/bin
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

export GO111MODULE=on

###############################################################################
###                                  Build                                  ###
###############################################################################
.PHONY: build install

all: install

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hid-noded		

build:
	go build $(BUILD_FLAGS) -o ./build/hid-noded ./cmd/hid-noded

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-build:
	sh ./scripts/protocgen.sh