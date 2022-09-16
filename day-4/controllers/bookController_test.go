package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookJSON = `{
		"id": 4,
		"title": "Hafalan Shalat Delisa",
		"author": "Tere Liye",
		"publisher": "Republika (Jakarta)"
	}`
)

func TestGetBookControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetBookControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetBookByIdControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		idParam              string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/books/:id",
			idParam:              "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.idParam)

		// Assertions
		if assert.NoError(t, GetBookByIdControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateBookControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		body                 io.Reader
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/books",
			body:                 strings.NewReader(bookJSON),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{",
		},
		{
			name:                 "Error Bad Request",
			path:                 "/books",
			body:                 strings.NewReader(`{"title":{},"author":{},"publisher":{}}`),
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", testCase.body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, CreateBookControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestUpdateBookControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		body                 io.Reader
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/books",
			body:                 strings.NewReader(bookJSON),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"book\":{",
		},
		{
			name:                 "Error Bad Request",
			path:                 "/books",
			body:                 strings.NewReader(`{"title":{},"author":{},"publisher":{}}`),
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", testCase.body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, UpdateBookControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestDeleteBookControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		idParam              string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/books/:id",
			idParam:              "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"book with ID",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.idParam)

		// Assertions
		if assert.NoError(t, DeleteBookControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
