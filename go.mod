module github.com/hypersign-protocol/hid-node

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.6
	github.com/cosmos/ibc-go v1.2.3
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/tendermint v0.34.20
	github.com/tendermint/tm-db v0.6.7
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.48.0
)

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/multiformats/go-multibase v0.0.3
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/viper v1.12.0
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
