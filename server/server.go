package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, &HealthCheckResponse{
		Message: "health!",
	})
}

func NewServer() (*echo.Echo, error) {
	verifyToken := os.Getenv("SLACK_VERIFY_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	if verifyToken == "" || botToken == "" {
		return nil, fmt.Errorf("slack env is empty")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/healthCheck", health)

	return e, nil
}
