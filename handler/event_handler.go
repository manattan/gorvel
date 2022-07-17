package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/gorvel/usecase"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type EventHandler struct {
	eu            *usecase.EventUseCase
	signingSecret string
}

func NewEventHandler(eu *usecase.EventUseCase, signingSecret string) *EventHandler {
	return &EventHandler{eu, signingSecret}
}

func (h *EventHandler) HandleEvent(c echo.Context) error {
	header := c.Request().Header

	verifier, err := slack.NewSecretsVerifier(header, h.signingSecret)
	if err != nil {
		return fmt.Errorf("could not verify as slack: %v", err)
	}

	bodyReader := io.TeeReader(c.Request().Body, &verifier)
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return fmt.Errorf("could not verify as slack: %v", err)
	}

	if err := verifier.Ensure(); err != nil {
		return fmt.Errorf("could not verify as slack: %v", err)
	}

	evt, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		return fmt.Errorf("could not parse event JSON: %v", err)
	}

	if evt.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			return fmt.Errorf("could not parse event response JSON: %v", err)
		}
		return c.JSON(http.StatusOK, &r)
	}

	h.eu.InvokeEvent(&evt)

	log.Println("Call event usecase with:", evt)
	return c.String(http.StatusOK, "handle event")
}
