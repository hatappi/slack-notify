// Package cmd handle command
package cmd

import (
	"github.com/hatappi/slack-notify/internal/slack"
)

// Option represents slack notice option
type Option func(*slack.PostMessageParams)

// WithAttachments set Attachments
func WithAttachments(a string) func(*slack.PostMessageParams) {
	return func(p *slack.PostMessageParams) {
		p.Attachments = &a
	}
}

// WithBlocks set Blocks
func WithBlocks(b string) func(*slack.PostMessageParams) {
	return func(p *slack.PostMessageParams) {
		p.Blocks = &b
	}
}

// SlackNotifyCmd represents command
type SlackNotifyCmd struct {
	Token   string
	Channel string
}

// Run executes command
func (snc *SlackNotifyCmd) Run(text string, opts ...Option) error {
	client := slack.NewClient()

	params := &slack.PostMessageParams{
		Token:   snc.Token,
		Channel: snc.Channel,
		Text:    text,
	}

	for _, opt := range opts {
		opt(params)
	}

	return client.Chat.PostMessage(params)
}
