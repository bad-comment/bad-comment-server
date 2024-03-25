package main

import (
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	ctl := &controller{}
	e.GET("/auth/sign_up", ctl.signUp)
	e.POST("/auth/sing_in", ctl.signIn)
	e.POST("/subjects", ctl.createSubject, jwtMiddleware(), userIdMiddleware)
	e.GET("/subjects/random", ctl.getSubjectByRandom, jwtMiddleware(), userIdMiddleware)
	e.POST("/subjects/:subjectId/comments", ctl.createSubjectComment, jwtMiddleware(), userIdMiddleware)
	e.GET("/trending/subjects", ctl.getTrendingSubjects)
	e.GET("/trending/kings", ctl.getTrendingKings)
	e.GET("/trending/queens", ctl.getTrendingQueens)
}
