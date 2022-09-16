package controllers

import (
	"alterra-agmc-day-4/config"
	"alterra-agmc-day-4/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com","password":"1234"}`
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	return e
}

func TestGetUserControllers(t *testing.T) {
	e := InitEcho()

	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetUserByIdControllers(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUserControllers(c)
	body, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	userCreated := responseBody["user"].(map[string]interface{})
	userCreatedID := userCreated["ID"]

	var testCases = []struct {
		name                 string
		path                 string
		idParam              string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"user\":{",
		},
		{
			name:                 "Error Not Found",
			path:                 "/users/:id",
			idParam:              "-1",
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "{\"message\":\"record not found\"}",
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
		if assert.NoError(t, GetUserByIdControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestCreateUserControllers(t *testing.T) {
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
			path:                 "/users",
			body:                 strings.NewReader(userJSON),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"user\":{",
		},
		{
			name:                 "Error Bad Request 1",
			path:                 "/users",
			body:                 strings.NewReader(`{"name":{},"email":{},"password":{}}`),
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
		{
			name:                 "Error Bad Request 2",
			path:                 "/users",
			body:                 nil,
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
		if assert.NoError(t, CreateUserControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestUpdateUserControllers(t *testing.T) {
	e := InitEcho()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUserControllers(c)
	body, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	userCreated := responseBody["user"].(map[string]interface{})
	userCreatedID := userCreated["ID"]

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	CreateUserControllers(c)
	body, _ = io.ReadAll(rec.Body)
	var responseBody2 map[string]interface{}
	json.Unmarshal(body, &responseBody2)
	userCreated2 := responseBody2["user"].(map[string]interface{})
	userCreatedID2 := userCreated2["ID"]

	var testCases = []struct {
		name                 string
		path                 string
		idParam              string
		userID               string
		body                 io.Reader
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID),
			userID:               fmt.Sprintf("%v", userCreatedID),
			body:                 strings.NewReader(userJSON),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"user\":{",
		},
		{
			name:                 "Error Not Found",
			path:                 "/users/:id",
			idParam:              "-1",
			userID:               "-1",
			expectStatus:         http.StatusNotFound,
			body:                 strings.NewReader(userJSON),
			expectBodyStartsWith: "{\"message\":\"record not found\"}",
		},
		{
			name:                 "Error Bad Request 1",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID),
			userID:               fmt.Sprintf("%v", userCreatedID),
			body:                 strings.NewReader(`{"name":{},"email":{},"password":{}}`),
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
		{
			name:                 "Error Forbidden",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID2),
			userID:               fmt.Sprintf("%v", userCreatedID),
			body:                 strings.NewReader(userJSON),
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: "{\"message\":\"You don't have permission\"}",
		},
	}

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPut, "/", testCase.body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		userId, _ := strconv.Atoi(testCase.userID)
		claims["userId"] = float64(userId)
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token.Valid = true
		c.Set("user", token)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.idParam)

		// Assertions
		if assert.NoError(t, UpdateUserControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestDeleteUserControllers(t *testing.T) {
	e := InitEcho()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUserControllers(c)
	body, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	userCreated := responseBody["user"].(map[string]interface{})
	userCreatedID := userCreated["ID"]

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	CreateUserControllers(c)
	body, _ = io.ReadAll(rec.Body)
	var responseBody2 map[string]interface{}
	json.Unmarshal(body, &responseBody2)
	userCreated2 := responseBody2["user"].(map[string]interface{})
	userCreatedID2 := userCreated2["ID"]

	var testCases = []struct {
		name                 string
		path                 string
		idParam              string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Success",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID),
			userID:               fmt.Sprintf("%v", userCreatedID),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\"}",
		},
		{
			name:                 "Error Not Found",
			path:                 "/users/:id",
			idParam:              "-1",
			userID:               "-1",
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "{\"message\":\"record not found\"}",
		},
		{
			name:                 "Error Forbidden",
			path:                 "/users/:id",
			idParam:              fmt.Sprintf("%v", userCreatedID2),
			userID:               fmt.Sprintf("%v", userCreatedID),
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: "{\"message\":\"You don't have permission\"}",
		},
	}

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		userId, _ := strconv.Atoi(testCase.userID)
		claims["userId"] = float64(userId)
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token.Valid = true
		c.Set("user", token)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.idParam)

		// Assertions
		if assert.NoError(t, DeleteUserControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}

func TestLoginUserControllers(t *testing.T) {
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
			path:                 "/login",
			body:                 strings.NewReader(userJSON),
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"token\":",
		},
		{
			name:                 "Error Bad Request 1",
			path:                 "/login",
			body:                 strings.NewReader(`{"email":{},"password":{}}`),
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
		{
			name:                 "Error Bad Request 2",
			path:                 "/login",
			body:                 nil,
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"message\":\"code=400",
		},
		{
			name:                 "Error Not Found",
			path:                 "/login",
			body:                 strings.NewReader(`{"email":"error@error.com","password":"-1"}`),
			expectStatus:         http.StatusNotFound,
			expectBodyStartsWith: "{\"message\":\"record not found\"}",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", testCase.body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, LoginUserControllers(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)

			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}

	}
}
