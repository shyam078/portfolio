package repository

import (
	"errors"
	"portifolio/internal/models"
)

type InMemory struct {
	Portifolio map[string]*models.Security
}

func NewInMemoryRepo() *InMemory {
	return &InMemory{
		Portifolio: make(map[string]*models.Security),
	}
}

func (r *InMemory) AddTrade(trade models.Trade) error {
	sec, exists := r.Portifolio[trade.TicketSymbol]

	if trade.TradeType == "Buy" {
		if exists {
			totalCost := sec.AverageBuyPrice*float64(sec.Quantity) + trade.Price*float64(trade.Quantity)
			totalQty := sec.Quantity + trade.Quantity
			sec.AverageBuyPrice = totalCost / float64(totalQty)
			sec.Quantity = totalQty
		} else {
			r.Portifolio[trade.TicketSymbol] = &models.Security{
				TicketSymbol:    trade.TicketSymbol,
				AverageBuyPrice: trade.Price,
				Quantity:        trade.Quantity,
			}
		}
	} else if trade.TradeType == "Sell" {
		if !exists || sec.Quantity < trade.Quantity {
			return errors.New("not enough shares to sell")
		}
		sec.Quantity -= trade.Quantity
	}

	return nil
}

func (r *InMemory) GetReturns(currentPrice float64) float64 {
	var total float64
	for _, sec := range r.Portifolio {
		gain := (currentPrice - sec.AverageBuyPrice) * float64(sec.Quantity)
		total += gain
	}
	return total
}
