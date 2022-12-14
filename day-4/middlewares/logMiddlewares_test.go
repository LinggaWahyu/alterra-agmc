package middlewares

import (
	"alterra-agmc-day-4/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogMiddlewares(t *testing.T) {
	e := echo.New()
	config.InitDB()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)

	LogMiddlewares(e)

	assert.NotEmpty(t, e)
	assert.NotNil(t, e)
}
