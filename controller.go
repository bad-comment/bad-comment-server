package main

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type (
	controller struct {
		db *gorm.DB
	}
)

func (ctl *controller) signUp(c echo.Context) error {
	var body SignUpDTO
	err := c.Bind(&body)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	// 注册
	user := User{
		Account:  body.Account,
		Password: body.Password,
	}
	result := ctl.db.Create(&user)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}

	// 登录
	res, err := loginService(user.Id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, res)
}

func getUserId(c echo.Context) (int64, bool) {
	if c.Get("userId") == nil {
		return 0, false
	}
	return int64(c.Get("userId").(int64)), true
}

func (ctl *controller) signIn(c echo.Context) error {
	var body SignInDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	var user User
	result := ctl.db.Where(&User{Account: body.Account, Password: body.Password}).First(&user)
	if result.Error != nil {
		return echo.ErrUnauthorized
	}

	// 登录
	res, err := loginService(user.Id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *controller) createSubject(c echo.Context) error {
	var userId, ok = getUserId(c)
	if !ok {
		return echo.ErrUnauthorized
	}

	var body CreateSubjectDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	var subject = Subject{
		Name:      body.Name,
		CreatedBy: userId,
	}
	result := ctl.db.Create(&subject)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, subject)
}

func (ctl *controller) getSubjectByRandom(c echo.Context) error {
	var userId, ok = getUserId(c)
	if !ok {
		return echo.ErrUnauthorized
	}

	var subject Subject
	result := ctl.db.Where("created_by <> ?", userId).Order("id desc").First(&subject)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, subject)
}

func (ctl *controller) createSubjectComment(c echo.Context) error {
	var userId, ok = getUserId(c)
	if !ok {
		return echo.ErrUnauthorized
	}

	var body CreateCommentDTO
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}
	subjectId, err := strconv.ParseInt(c.Param("subjectId"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	var comment = SubjectComment{
		UserId:    userId,
		SubjectId: subjectId,
		Score:     body.Score,
	}
	result := ctl.db.Create(&comment)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, comment)
}

type SubjectHot struct {
	SubjectId int64 `json:"subject_id"`
	Score     int64 `json:"score"`
}

func (ctl *controller) getTrendingKings(c echo.Context) error {
	var userScores []UserScore
	result := ctl.db.Order("score desc").Limit(100).Find(&userScores)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, userScores)
}

func (ctl *controller) getTrendingQueens(c echo.Context) error {
	var userScores []UserScore
	result := ctl.db.Order("score desc").Limit(100).Find(&userScores)
	if result.Error != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, userScores)
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
