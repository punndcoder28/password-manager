package vault

type VaultFile struct {
	Version    int    `json:"version"`
	Salt       []byte `json:"salt"`
	Nonce      []byte `json:"nonce"`
	CypherText []byte `json:"cypher_text"`
}
