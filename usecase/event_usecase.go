package usecase

import (
	"fmt"
	"math/rand"

	"github.com/manattan/gorvel/repository"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var randomMessages = []string{
	"Did you need me?",
	"What's up?",
	"I'm quite tired...",
	"I don't wanna work any more",
}

type EventUseCase struct {
	slackRepository *repository.SlackRepository
}

func NewEventUseCase(slackRepository *repository.SlackRepository) *EventUseCase {
	return &EventUseCase{slackRepository}
}

func (eu *EventUseCase) InvokeEvent(evt *slackevents.EventsAPIEvent) error {
	switch evt := evt.InnerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		return eu.slackRepository.PostMessage(evt.Channel, slack.MsgOptionText(randomMessages[rand.Intn(len(randomMessages))], false))
	default:
		return fmt.Errorf("this event is not supported : %v", evt)
	}
}
