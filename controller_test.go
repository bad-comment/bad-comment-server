package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	const body = `{"account":"zhuzhu","password":"123456"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/sign_up", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.signUp(c)) {
		var res AuthToken
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "hahahahhahahaah", res.Token)
	}
}

func TestSignIn(t *testing.T) {
	const body = `{"account":"zhuzhu","password":"123456"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/sign_in", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.signIn(c)) {
		var res AuthToken
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "hahahahhahahaah", res.Token)
	}
}

func TestGetSubjects(t *testing.T) {
	const expected = `[{"id":1,"name":"铅笔","created_by":1}]` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/subjects", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.getSubjects(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetSubjectByRandom(t *testing.T) {
	const expected = `{"id":1,"name":"铅笔","created_by":1}` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/subject/random", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.getSubjectByRandom(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestCreateSubjectComment(t *testing.T) {
	const body = `{"score":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/subjects/1/comments", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.createSubjectComment(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, 0, rec.Body.Len())
	}
}

func TestGetTrendingSubjects(t *testing.T) {
	const expected = `[{"subject_id":1,"score":120}]` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/trending/subjects", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.getTrendingSubjects(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetTrendingKings(t *testing.T) {
	const expected = `[{"user_id":1,"score":130}]` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.getTrendingKings(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetTrendingQueens(t *testing.T) {
	const expected = `[{"user_id":1,"score":110}]` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := &controller{}

	if assert.NoError(t, ctl.getTrendingQueens(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
