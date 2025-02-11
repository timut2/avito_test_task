package models

type InfoResponse struct {
	Coins int
	// inventory   Inventory
	// coinHistory CoinHistory
}

type ErrorResponse struct {
	Errors error
}
