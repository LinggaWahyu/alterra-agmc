package middlewares

import (
	"alterra-agmc-day-4/config"
	"alterra-agmc-day-4/lib/database"
	"alterra-agmc-day-4/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	config.InitDB()
	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := database.CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	token, err := CreateToken(int(userCreated.ID))
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestExtractTokenUserId(t *testing.T) {
	e := echo.New()
	config.InitDB()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}

	userCreated, err := database.CreateUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userCreated)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = float64(userCreated.ID)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Valid = true
	c.Set("user", token)

	userIDSuccess := ExtractTokenUserId(c)
	assert.NotEmpty(t, userIDSuccess)
	assert.NotEqual(t, userIDSuccess, 0)
	assert.Equal(t, uint(userIDSuccess), userCreated.ID)

	token.Valid = false
	c.Set("user", token)

	userIDError := ExtractTokenUserId(c)
	assert.Empty(t, userIDError)
	assert.Equal(t, userIDError, 0)
	assert.NotEqual(t, uint(userIDError), userCreated.ID)
}
