package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	verifyToken string
}

func NewEventHandler (verifyToken string) *EventHandler {
	return &EventHandler{verifyToken: verifyToken}
}

func (h *EventHandler) HandleEvent (c echo.Context) error {
	return c.String(http.StatusOK, "handle event")
}