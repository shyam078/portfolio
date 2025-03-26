package repository

import "portifolio/internal/models"

type TradeReporitory interface {
	RecordTrade(trade models.Trade) error
}
