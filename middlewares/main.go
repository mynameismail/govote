package middlewares

import (
	"govote/models"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var AuthAdmin = middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
	if username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD") {
		return true, nil
	}

	return false, nil
})

var AuthVoter = middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	KeyLookup: "header:voter-code",
	Validator: func(key string, c echo.Context) (bool, error) {
		var voter models.Voter
		notfound := models.Conn.Where("code = ? AND status = 'logged'", key).First(&voter).RecordNotFound()
		if notfound {
			return false, nil
		}

		c.Set("voter", voter)

		return true, nil
	},
})
