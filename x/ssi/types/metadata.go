package types

import (
	"encoding/base64"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func CreateNewMetadata(ctx sdk.Context) Metadata {
	return Metadata{
		VersionId:   base64.StdEncoding.EncodeToString(tmhash.Sum([]byte(ctx.TxBytes()))),
		Deactivated: false,
		Created:     ctx.BlockTime().Format(time.RFC3339), //Ref: https://www.w3.org/TR/did-core/#did-document-metadata
		Updated:     ctx.BlockTime().Format(time.RFC3339),
	}
}
