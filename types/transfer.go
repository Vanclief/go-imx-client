package types

type Transfer struct {
	TransactionID int    `json:"transaction_id"`
	Status        string `json:"status"`
	User          string `json:"user"`
	Receiver      string `json:"receiver"`
	Token         Token  `json:"token"`
	Timestamp     string `json:"timestamp"`
}

type Token struct {
	Type string    `json:"type"`
	Data TokenData `json:"data"`
}

type TokenData struct {
	TokenID      string `json:"token_id"`
	ID           string `json:"id"`
	TokenAddress string `json:"token_address"`
	Decimals     int    `json:"decimals"`
	Quantity     string `json:"quantity"`
}
