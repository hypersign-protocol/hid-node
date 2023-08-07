package bindings

// SsiContractQuery contains custom queries for x/ssi chain
type SsiContractQuery struct {
	DidDocumentExists *QueryDidDocumentExistsRequest `json:"did_document_exists,omitempty"`
}

type QueryDidDocumentExistsRequest struct {
	DidId string `json:"did_id,omitempty"`
}

type QueryDidDocumentExistsResponse struct {
	Result bool `json:"result"`
}
