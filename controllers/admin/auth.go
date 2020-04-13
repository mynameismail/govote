package admin

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    type Reqbody struct {
        Username string `json:"username" form:"username"`
        Password string `json:"password" form:"password"`
    }

    reqbody := new(Reqbody)
    if c.Bind(reqbody) != nil {
        return c.String(400, "Sorry! Bad request")
    }
    
    if reqbody.Username == os.Getenv("ADMIN_USERNAME") && reqbody.Password == os.Getenv("ADMIN_PASSWORD") {
        return c.String(200, "Correct login")
    }

    return c.String(401, "Wrong login")
}
