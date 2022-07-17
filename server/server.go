package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manattan/gorvel/handler"
	"github.com/manattan/gorvel/repository"
	"github.com/manattan/gorvel/usecase"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func health(c echo.Context) error {
	var r *slackevents.ChallengeResponse
	log.Println(r)
	return c.JSON(http.StatusOK, &r)
}

func NewServer() (*echo.Echo, error) {
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	if signingSecret == "" || botToken == "" {
		return nil, fmt.Errorf("slack env is empty")
	}

	client := slack.New(botToken)
	_, err := client.AuthTest()
	if err != nil {
		return nil, err
	}

	sr := repository.NewSlackRepository(client)

	eu := usecase.NewEventUseCase(sr)

	eventHandler := handler.NewEventHandler(eu, signingSecret)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/healthCheck", health)
	e.POST("/events", eventHandler.HandleEvent)

	return e, nil
}
