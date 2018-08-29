package server

import (
	"github.com/labstack/echo"
)

func CreateEndpoint() *echo.Echo {
	e := echo.New()

	e.GET("/report", handleReportState)
	e.GET("/state", handleGetState)

	return e
}
