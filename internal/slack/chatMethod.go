package slack

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/morikuni/failure"

	"github.com/hatappi/slack-notify/internal/errors"
)

type ChatMethod struct {
	apiClient APIClient
}

func newChatMethod(apiClient APIClient) *ChatMethod {
	return &ChatMethod{
		apiClient: apiClient,
	}
}

type PostMessageParams struct {
	Token          string  `url:"token" validate:"required"`
	Channel        string  `url:"channel" validate:"required"`
	Text           string  `url:"text"`
	AsUser         *bool   `url:"as_user,omitempty"`
	Attachments    *string `url:"attachments,omitempty"`
	Blocks         *string `url:"blocks,omitempty"`
	IconEmoji      *string `url:"icon_emoji,omitempty"`
	IconURL        *string `url:"icon_url,omitempty"`
	LinkNames      *bool   `url:"link_names,omitempty"`
	Mrkdwn         *bool   `url:"mrkdwn,omitempty"`
	Parse          *string `url:"parse,omitempty"`
	ReplyBroadcase *bool   `url:"reply_broadcast,omitempty"`
	ThreadTS       *string `url:"thread_ts,omitempty"`
	UnfurlLinks    *bool   `url:"unfurl_links,omitempty"`
	UnfurlMedia    *bool   `url:"unfurl_media,omitempty"`
	Username       *string `url:"username,omitempty"`
}

// PostMessage post message to Slack channel
// https://api.slack.com/methods/chat.postMessage
func (c *ChatMethod) PostMessage(params PostMessageParams) error {
	validate := validator.New()
	err := validate.Struct(params)
	if err != nil {
		return failure.Wrap(err)
	}

	body, _, err := c.apiClient.Post("chat.postMessage", params)
	if err != nil {
		return failure.Translate(err, errors.FailedRequest)
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return failure.Wrap(err)
	}

	if isOk, ok := m["ok"]; ok {
		if isOk.(bool) {
			return nil
		}

		return failure.New(errors.FailedRequest, failure.Messagef("failed to request. reason is %s", m["error"]))
	}

	return nil
}
