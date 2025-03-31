package router

import (
	"portifolio/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, c *controller.TradeController) {
	r.POST("/trade", c.AddTrade)
	r.GET("/returns", c.GetRetuns)
	r.GET("/holdings", c.GetHoldings)
}
