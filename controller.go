package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SignUpDTO struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AuthToken struct {
	Token       string `json:"token"`
	ExpiredTime int    `json:"expired_time"`
}

type (
	controller struct {
	}
)

func (ctl *controller) signUp(c echo.Context) error {
	var body SignUpDTO
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	res := AuthToken{
		Token:       "hahahahhahahaah",
		ExpiredTime: 1000,
	}
	return c.JSON(http.StatusCreated, res)
}

func (ctl *controller) signIn(c echo.Context) error {
	var body SignUpDTO
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	res := AuthToken{
		Token:       "hahahahhahahaah",
		ExpiredTime: 1000,
	}
	return c.JSON(http.StatusOK, res)
}

type Subject struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedBy int64  `json:"created_by"`
}

func (ctl *controller) getSubjects(c echo.Context) error {
	var res = []Subject{
		{
			Id:        1,
			Name:      "铅笔",
			CreatedBy: 1,
		},
	}
	return c.JSON(http.StatusOK, res)
}

func (ctl *controller) getSubjectByRandom(c echo.Context) error {
	var res = Subject{
		Id:        1,
		Name:      "铅笔",
		CreatedBy: 1,
	}
	return c.JSON(http.StatusOK, res)
}

type SubjectComment struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"user_id"`
	SubjectId int64 `json:"subject_id"`
	Score     int8  `json:"score"`
}

type CreateCommentDTO struct {
	Score int8 `json:"score"`
}

func (ctl *controller) createSubjectComment(c echo.Context) error {
	var body CreateCommentDTO
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.NoContent(http.StatusNoContent)
}

type UserScore struct {
	UserId int64 `json:"user_id"`
	Score  int64 `json:"score"`
}

type SubjectHot struct {
	SubjectId int64 `json:"subject_id"`
	Score     int64 `json:"score"`
}

func (ctl *controller) getTrendingSubjects(c echo.Context) error {
	var res = []SubjectHot{
		{
			SubjectId: 1,
			Score:     120,
		},
	}
	return c.JSON(http.StatusOK, res)
}

func (ctl *controller) getTrendingKings(c echo.Context) error {
	var res = []UserScore{
		{
			UserId: 1,
			Score:  130,
		},
	}
	return c.JSON(http.StatusOK, res)
}

func (ctl *controller) getTrendingQueens(c echo.Context) error {
	var res = []UserScore{
		{
			UserId: 1,
			Score:  110,
		},
	}
	return c.JSON(http.StatusOK, res)
}
