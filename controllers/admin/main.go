package admin

import (
	"govote/models"

	"github.com/labstack/echo/v4"
)

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

func ListVoters(c echo.Context) error {
	type Response struct {
		Message string         `json:"message"`
		Data    []models.Voter `json:"data"`
	}

	var voters []models.Voter
	models.Conn.Find(&voters)

	return c.JSON(200, &Response{
		Message: "Success",
		Data:    voters,
	})
}

func CreateVoting(c echo.Context) error {
	type Reqbody struct {
		Names          []string `json:"names" form:"names"`
		NumberOfVoters int      `json:"number_of_voters" form:"number_of_voters"`
		Prefix         string   `json:"prefix" form:"prefix"`
	}

	reqbody := new(Reqbody)
	err := c.Bind(reqbody)
	if err != nil || len(reqbody.Names) < 1 || reqbody.NumberOfVoters < 1 || reqbody.Prefix == "" {
		return c.String(400, "Sorry! Bad request")
	}

	for _, name := range reqbody.Names {
		if name != "" {
			vote := models.Vote{Name: name}
			models.Conn.Create(&vote)
		}
	}

	prefix := reqbody.Prefix
	for i := 0; i < reqbody.NumberOfVoters; i++ {
		code := prefix + models.RandStringBytes(10)
		voter := models.Voter{Code: code}
		models.Conn.Create(&voter)
	}

	return c.String(200, "Success")
}

func ResetVoting(c echo.Context) error {
	var votes models.Vote
	var voters models.Voter
	models.Conn.Delete(&votes)
	models.Conn.Delete(&voters)

	return c.String(200, "Success")
}
