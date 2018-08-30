package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func handleAuthorizeView(c echo.Context) error {

	// clientId := c.QueryParam("client_id")
	redirectUrl := c.QueryParam("redirect_uri")
	state := c.QueryParam("state")
	// responseType := c.QueryParam("response_type")

	return c.Render(http.StatusOK, "authorize.html", map[string]string{
		"redirectUrl": redirectUrl,
		"state":       state,
	})
}

func handleAuthorize(c echo.Context) error {

	redirectUrl := c.QueryParam("redirect_uri")
	state := c.QueryParam("state")
	token := GenerateID(32)
	// code := GenerateID(16)

	url := redirectUrl + "#state=" + state + "&access_token=" + token + "&token_type=bearer"

	return c.Redirect(http.StatusMovedPermanently, url)
}
