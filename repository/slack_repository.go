package repository

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackRepository struct {
	client *slack.Client
}

func NewSlackRepository(client *slack.Client) *SlackRepository {
	return &SlackRepository{client}
}

func (sr SlackRepository) PostMessage(channelID string, msg ...slack.MsgOption) error {
	_, _, err := sr.client.PostMessage(channelID, msg...)
	if err != nil {
		return fmt.Errorf("could not parse slash command JSON: %v", err)
	}
	return nil
}
