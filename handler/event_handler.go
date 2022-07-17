package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/gorvel/usecase"
	"github.com/slack-go/slack/slackevents"
)

type EventHandler struct {
	eu *usecase.EventUseCase
	verifyToken string
}

func NewEventHandler (eu *usecase.EventUseCase, verifyToken string) *EventHandler {
	return &EventHandler{eu, verifyToken}
}

func (h *EventHandler) HandleEvent (c echo.Context) error {
	defer c.Request().Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	body := buf.String()

	evt, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: h.verifyToken}))
	if err != nil {
		return fmt.Errorf("could not parse event JSON: %v", err)
	}

	if evt.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			return fmt.Errorf("could not parse event response JSON: %v", err)
		}
	}

	h.eu.InvokeEvent(&evt)

	log.Println("Call event usecase with:", evt)
	return c.String(http.StatusOK, "handle event")
}