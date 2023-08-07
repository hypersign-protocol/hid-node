package bindings

type SsiContractMsg struct {
	SetBlockchainAccountId *SetBlockchainAccountIdStruct `json:"set_blockchain_account_id,omitempty"`
}

type SetBlockchainAccountIdStruct struct {
	BlockchainAccountId string `json:"blockchain_account_id,omitempty"`
}