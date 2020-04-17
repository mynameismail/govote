package voter

import (
	"govote/models"

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
		return c.String(401, "Invalid login")
	}

	voter.Status = "logged"
	models.Conn.Save(&voter)

	return c.String(200, "Valid login")
}

func ListVotes(c echo.Context) error {
	type Response struct {
		Message string        `json:"message"`
		Data    []models.Vote `json:"data"`
	}

	var votes []models.Vote
	models.Conn.Find(&votes)

	return c.JSON(200, &Response{
		Message: "Success",
		Data:    votes,
	})
}

func Vote(c echo.Context) error {
	type Reqbody struct {
		SelectedID int `json:"selected_id" form:"selected_id"`
	}

	reqbody := new(Reqbody)
	if c.Bind(reqbody) != nil {
		return c.String(400, "Sorry! Bad request")
	}

	var vote models.Vote
	notfound := models.Conn.First(&vote, reqbody.SelectedID).RecordNotFound()
	if notfound {
		return c.String(404, "Invalid selection")
	}

	vote.Votes = vote.Votes + 1
	models.Conn.Save(&vote)

	voter := c.Get("voter").(models.Voter)
	voter.Status = "voted"
	models.Conn.Save(&voter)

	return c.String(200, "Success")
}
