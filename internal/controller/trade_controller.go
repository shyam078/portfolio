package controller

import (
	"net/http"
	"portifolio/internal/models"
	"portifolio/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TradeController struct {
	Service *service.PortifolioService
}

func NewTradeController(s *service.PortifolioService) *TradeController {
	return &TradeController{
		Service: s,
	}
}

func (tc *TradeController) AddTrade(c *gin.Context) {
	var trade models.Trade

	if err := c.ShouldBindBodyWithJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := tc.Service.AddTradeService(trade)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Trade Added",
		})
	}
}

func (tc *TradeController) GetRetuns(c *gin.Context) {
	priceStr := c.Query("price")
	if priceStr == "" {
		priceStr = "100"
	}
	price, _ := strconv.ParseFloat(priceStr, 64)

	returns := tc.Service.GetReturns(price)

	c.JSON(http.StatusOK, gin.H{
		"returns": returns,
	})
}
