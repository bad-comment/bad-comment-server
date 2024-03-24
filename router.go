package main

import (
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	ctl := &controller{}

	e.GET("/auth/sign_up", ctl.signUp)
	e.POST("/auth/sing_in", ctl.signIn)
	e.GET("/subjects", ctl.getSubjects)
	e.GET("/subjects/random", ctl.getSubjectByRandom)
	e.POST("/subjects/:id/comments", ctl.createSubjectComment)
	e.GET("/trending/subjects", ctl.getTrendingSubjects)
	e.GET("/trending/kings", ctl.getTrendingKings)
	e.GET("/trending/queens", ctl.getTrendingQueens)
}
