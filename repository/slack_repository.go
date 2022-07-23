package repository

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackRepository interface {
	PostMessage(channelID string, msg ...slack.MsgOption) error
}

var _ SlackRepository = &slackRepository{}

type slackRepository struct {
	client *slack.Client
}


// TODO: Clientのinterfaceがないので、どういう感じにすればいいか考える
func NewSlackRepository(client *slack.Client) *slackRepository {
	return &slackRepository{client}
}

func (sr *slackRepository) PostMessage(channelID string, msg ...slack.MsgOption) error {
	_, _, err := sr.client.PostMessage(channelID, msg...)
	if err != nil {
		return fmt.Errorf("could not parse slash command JSON: %v", err)
	}
	return nil
}