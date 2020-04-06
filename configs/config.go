package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type SlackNotifyCmdConfig struct {
	Token       string `envconfig:"SLACK_NOTIFY_TOKEN"`
	Channel     string `envconfig:"SLACK_NOTIFY_CHANNEL"`
	Text        string `envconfig:"SLACK_NOTIFY_TEXT"`
	Attachments string `envconfig:"SLACK_NOTIFY_ATTACHMENTS"`
	Blocks      string `envconfig:"SLACK_NOTIFY_BLOCKS"`
}

func LoadSlackNotifyCmdConfig() (*SlackNotifyCmdConfig, error) {
	var c SlackNotifyCmdConfig
	err := envconfig.Process("SLACK_NOTIFY", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
