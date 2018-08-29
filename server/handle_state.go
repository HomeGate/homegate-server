package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func handleReportState(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func handleGetState(c echo.Context) error {

	return c.JSON(http.StatusOK, State{
		Properties: []Property{
			{
				Action: CameraStream,
				Value:  "https://camerastream.homegate.cc/000-00-000",
			},
			{
				Action: OnOff,
				Value:  "ON",
			},
		},
	})
}
