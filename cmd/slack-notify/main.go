package main

import (
	"flag"
	"log"

	"github.com/hatappi/slack-notify/configs"
	"github.com/hatappi/slack-notify/internal/cmd"
)

var (
	token       = flag.String("token", "", "SlackBot token")
	channel     = flag.String("channel", "", "notify channel on Slack")
	text        = flag.String("text", "", "notify text")
	attachments = flag.String("attachments", "", "notify text")
	blocks      = flag.String("blocks", "", "notify text")
)

func main() {
	flag.Parse()

	config, err := configs.LoadSlackNotifyCmdConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := &cmd.SlackNotifyCmd{
		Token:   config.Token,
		Channel: config.Channel,
	}

	// merge command line args
	if *token != "" {
		c.Token = *token
	}
	if *channel != "" {
		c.Channel = *channel
	}
	if *text != "" {
		config.Text = *text
	}
	if *attachments != "" {
		config.Attachments = *attachments
	}
	if *blocks != "" {
		config.Blocks = *blocks
	}

	opts := []cmd.Option{}

	if config.Attachments != "" {
		opts = append(opts, cmd.WithAttachments(config.Attachments))
	}

	if config.Blocks != "" {
		opts = append(opts, cmd.WithBlocks(config.Blocks))
	}

	c.Run(config.Text, opts...)
}
