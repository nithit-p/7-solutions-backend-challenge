package app

import (
	"log"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app/beef"

	"github.com/labstack/echo/v4"
)

type HttpHandlersConfig struct {
	BeefHandler beef.BeefHttpServer
}

func StartHttpServer(address string, handlersConfig HttpHandlersConfig) (*echo.Echo, error) {
	e := echo.New()

	e.POST("/beef/summary", handlersConfig.BeefHandler.GetBeefSummary)

	e.Logger.Fatal(e.Start(address))
	log.Println("Http server started at " + address)

	return e, nil
}
