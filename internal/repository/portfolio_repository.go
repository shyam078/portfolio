package repository

import "portifolio/internal/models"

type PortifolioRepository interface {
	UpdatePortifolio(trade models.Trade) error
	FetchPortifolio() ([]models.Security, error)
}
