package models

type Security struct {
	TicketSymbol    string  `json:"ticker"`
	AverageBuyPrice float64 `json:"avg_buy_price"`
	Quantity        int     `json:"quantity"`
}
