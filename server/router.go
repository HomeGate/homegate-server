package server

import (
	"github.com/labstack/echo"
	"html/template"
)

func CreateEndpoint() *echo.Echo {

	e := echo.New()

	renderer := &Renderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = renderer

	e.Static("/static", "server/static")

	e.POST("/bridge/google", handleGoogleRequests)
	e.POST("/bridge/alexa", handleGoogleRequests)

	e.GET("/authorize", handleAuthorizeView)
	e.GET("/authorize/continue", handleAuthorize)

	return e
}
