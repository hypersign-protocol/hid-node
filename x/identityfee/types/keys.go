package types

const (
	ModuleName = "identityfee"

	QuerierRoute = ModuleName
)

var (
	ParamStoreKeyCreateDidFee = []byte("CreateDidFee")
	ParamStoreKeyUpdateDidFee = []byte("UpdateDidFee")
	ParamStoreKeyDeactivateDidFee = []byte("DeactivateDidFee")
	ParamStoreKeyCreateSchemaFee = []byte("CreateSchemaFee")
	ParamStoreKeyRegisterCredentialStatusFee = []byte("RegisterCredentialStatusFee")
)