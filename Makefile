#!/usr/bin/make -f

VERSION := $(shell git describe --tags --abbrev=0)
COMMIT := $(shell git rev-parse --short HEAD)

BUILD_DIR ?= $(CURDIR)/build
HIDNODE_CMD_DIR := $(CURDIR)/cmd/hid-noded
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

GOBIN = $(shell go env GOPATH)/bin
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

SDK_VERSION := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's:.* ::')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=hid-node \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

BUILD_FLAGS := -ldflags '$(ldflags)'

export GO111MODULE=on

###############################################################################
###                                  Build                                  ###
###############################################################################
.PHONY: build install 

all: proto-gen-go proto-gen-swagger build

go-version-check:
ifneq ($(GO_MINOR_VERSION),19)
	@echo "ERROR: Go version 1.19 is required to build hid-noded binary"
	exit 1
endif

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

install: go-version-check go.sum
	go install -mod=readonly $(BUILD_FLAGS) $(HIDNODE_CMD_DIR)	

build: go-version-check
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILD_DIR)/hid-noded $(HIDNODE_CMD_DIR)

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-gen-go:
	@echo "Generating golang code from protobuf"
	./scripts/protocgen-go.sh

proto-gen-swagger:
	@echo "Generating swagger docs"
	./scripts/protoc-swagger-gen.sh

proto-gen-ts:
	@echo "Generating typescript code from protobuf"
	./scripts/protocgen-ts.sh

###############################################################################
###                                  Docker                                 ###
###############################################################################
DOCKER_IMAGE_NAME := hid-node-image

docker-all: docker-build docker-run

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run --rm -d \
	-p 26657:26657 -p 1317:1317 -p 26656:26656 -p 9090:9090 \
	--name hid-node-container \
	$(DOCKER_IMAGE_NAME) start
