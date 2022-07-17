package usecase

import (
	"github.com/manattan/gorvel/repository"
	"github.com/slack-go/slack/slackevents"
)

type EventUseCase struct {
	slackRepository *repository.SlackRepository
}

func NewEventUseCase (slackRepository *repository.SlackRepository) *EventUseCase {
	return &EventUseCase{slackRepository}
}

func (eu *EventUseCase) InvokeEvent(evt *slackevents.EventsAPIEvent) error {
	return nil
}