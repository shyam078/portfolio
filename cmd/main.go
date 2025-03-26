package main

import (
	"portifolio/internal/controller"
	"portifolio/internal/repository"
	"portifolio/internal/service"
	"portifolio/pkg/utils/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewInMemoryRepo()

	svc := service.NewPortifolioService(repo)

	ctl := controller.NewTradeController(svc)

	router.RegisterRoute(r, ctl)

	r.Run(":8090")
}
