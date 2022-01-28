package types

const (
	// ModuleName defines the module name
	ModuleName = "did"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_did"
)

const (
	DidKey      = "Did-value-"
	DidCountKey = "Did-count-"
)

const (
	SchemaKey      = "Schema-value-"
	SchemaCompleteKey = "Schema-complete-"
	SchemaCountKey = "Schema-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
