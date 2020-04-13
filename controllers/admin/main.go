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

func CreateVotes(c echo.Context) error {
	type Reqbody struct {
		Names []string `json:"names" form:"names"`
	}

	reqbody := new(Reqbody)
	if c.Bind(reqbody) != nil {
		return c.String(400, "Sorry! Bad request")
	}

	for _, name := range reqbody.Names {
		vote := models.Vote{Name: name}
		models.Conn.Create(&vote)
	}

	return c.String(200, "Success")
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

func GenerateVoters(c echo.Context) error {
	type Reqbody struct {
		NumberOfVoters int    `json:"number_of_voters" form:"number_of_voters"`
		Prefix         string `json:"prefix" form:"prefix"`
	}

	reqbody := new(Reqbody)
	if c.Bind(reqbody) != nil {
		return c.String(400, "Sorry! Bad request")
	}

	prefix := reqbody.Prefix
	for i := 0; i < reqbody.NumberOfVoters; i++ {
		code := prefix + models.RandStringBytes(10)
		voter := models.Voter{Code: code}
		models.Conn.Create(&voter)
	}

	return c.String(200, "Success")
}
