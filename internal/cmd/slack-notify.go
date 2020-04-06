package cmd

import (
	"github.com/hatappi/slack-notify/internal/slack"
)

type Option func(*slack.PostMessageParams)

func WithAttachments(a string) func(*slack.PostMessageParams) {
	return func(p *slack.PostMessageParams) {
		p.Attachments = &a
	}
}

func WithBlocks(b string) func(*slack.PostMessageParams) {
	return func(p *slack.PostMessageParams) {
		p.Blocks = &b
	}
}

type SlackNotifyCmd struct {
	Token   string
	Channel string
}

func (snc *SlackNotifyCmd) Run(text string, opts ...Option) {
	client := slack.NewClient()

	params := slack.PostMessageParams{
		Token:   snc.Token,
		Channel: snc.Channel,
		Text:    text,
	}

	for _, opt := range opts {
		opt(&params)
	}

	client.Chat.PostMessage(params)
}
