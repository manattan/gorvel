package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/gorvel/usecase"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type EventHandler struct {
	eu            usecase.EventUseCase
	signingSecret string
}

func NewEventHandler(eu usecase.EventUseCase, signingSecret string) *EventHandler {
	return &EventHandler{eu, signingSecret}
}

func (h *EventHandler) HandleEvent(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Printf("%v", body)
		return fmt.Errorf("read all error: %v", err)
	}

	sv, err := slack.NewSecretsVerifier(c.Request().Header, h.signingSecret)
	if err != nil {
		fmt.Printf("%v", sv)
		return fmt.Errorf("secrets verify error: %v", err)
	}

	if _, err := sv.Write(body); err != nil {
		fmt.Printf("%v", sv)
		return fmt.Errorf("write error: %v", err)
	}

	if err := sv.Ensure(); err != nil {
		fmt.Printf("%v", sv)
		return fmt.Errorf("ensure error: %v", err)
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
