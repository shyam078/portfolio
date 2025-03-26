package models

type Trade struct {
	TicketSymbol string  `json:"ticker"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	TradeType    string  `json:"type"`
}
