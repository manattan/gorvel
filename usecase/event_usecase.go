package usecase

import (
	"fmt"
	"math/rand"

	"github.com/manattan/gorvel/repository"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var randomMessages = []string{
	"Hello, Mogi. Calm Down. You wanna drink Shiso juice. ",
	"What's up?",
	"I wanna be an Iron Man. And you?",
	"Sleepy.....",
}

type EventUseCase interface {
	InvokeEvent(evt *slackevents.EventsAPIEvent) error
}

var _ EventUseCase = &eventUseCase{}

type eventUseCase struct {
	slackRepository repository.SlackRepository
}

func NewEventUseCase(slackRepository repository.SlackRepository) *eventUseCase {
	return &eventUseCase{slackRepository}
}

func (eu *eventUseCase) InvokeEvent(evt *slackevents.EventsAPIEvent) error {
	switch evt := evt.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		return eu.slackRepository.PostMessage(evt.Channel, slack.MsgOptionText(randomMessages[rand.Intn(len(randomMessages))], false))
	default:
		return fmt.Errorf("this event is not supported : %v", evt)
	}
}
