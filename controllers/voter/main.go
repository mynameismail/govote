package voter

import (
	"govote/models"
	"log"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    type Reqbody struct {
        Code string `json:"code" form:"code"`
    }

    reqbody := new(Reqbody)
    if c.Bind(reqbody) != nil {
        return c.String(400, "Sorry! Bad request")
    }

    var voter models.Voter
    notfound := models.Conn.Where("code = ? AND status = 'init'", reqbody.Code).First(&voter).RecordNotFound()
    if notfound {
        return c.String(401, "Wrong login")
    }

    voter.Status = "logged"
    models.Conn.Save(&voter)
    return c.String(200, "Correct login")
}

func Vote(c echo.Context) error {
    log.Println(c.Get("voter"))

    type Reqbody struct {
        ID string `json:"id" form:"id"`
    }

    reqbody := new(Reqbody)
    if c.Bind(reqbody) != nil {
        return c.String(400, "Sorry! Bad request")
    }

    var voter models.Voter
    notfound := models.Conn.First(&voter, 1).RecordNotFound()
    if notfound {
        return c.String(401, "Wrong login")
    }

    log.Println(voter)

    // voter.Status = "voted"
    // models.Conn.Save(&voter)
    return c.String(200, "Success")
}
