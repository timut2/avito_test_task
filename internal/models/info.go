package models

type InfoResponse struct {
	Coins       int         `json:"coins"`
	Inventory   []Item      `json:"invetory"`
	CoinHistory CoinHistory `json:"coinHistory"`
}

type Item struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistory struct {
	Received []Transaction `json:"received"`
	Send     []Transaction `json:"sent"`
}

type Transaction struct {
	FromUser string `json:"fromUser ,omitempty"`
	ToUser   string `json:"toUser ,omitempty"`
	Amount   int    `json:"amount"`
}
