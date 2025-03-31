package service

import (
	"fmt"
	"portifolio/internal/models"
	"portifolio/internal/repository"
)

type PortifolioService struct {
	PortifolioUseCase *repository.InMemory
}

type PortfolioServiceInterface interface {
	AddTradeService(trade models.Trade) error

	GetHoldingsService() []map[string]interface{}

	GetReturns(currentPrice float64) float64
}

func NewPortifolioService(repo *repository.InMemory) *PortifolioService {
	return &PortifolioService{
		PortifolioUseCase: repo,
	}
}

func (p *PortifolioService) AddTradeService(trade models.Trade) error {
	err := p.PortifolioUseCase.AddTrade(trade)
	if err != nil {
		fmt.Println("Error while adding hthe trade")
	}
	return err
}

func (p *PortifolioService) GetReturns(currentPrice float64) float64 {
	return p.GetReturns(currentPrice) //Note: Instead of calling the repository method, I called the service, which is why we got the error
}
func (p *PortifolioService) GetHoldingsService() []map[string]interface{} {
	return p.PortifolioUseCase.GetHoldings()
}