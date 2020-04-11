# SlackNotify
slackNotify is simple command line tool to notify message to Slack.

## Installation

```sh
$ go get github.com/hatappi/slack-notify/cmd/slack-notify
```

or 

https://github.com/hatappi/slack-notify/releases/latest

## Usage

```sh
export SLACK_NOTIFY_TOKEN=xoxb-xxxxxxxxx
export SLACK_NOTIFY_CHANNEL=CXXXXXXXX
```

### Local

```
// e.g.1
$ slack-notify -text "test"

// e.g.2
$ slack-notify -attachments '[{"pretext": "pre-hello", "color": "#00CCFF","text": "hello"}]'
```

### Docker

```sh
// e.g.1 
docker run -it \
	-e SLACK_NOTIFY_TOKEN \
	-e SLACK_NOTIFY_CHANNEL \
	hatappi/slack-notify \
	-text test
	
// e.g.2
docker run -it \
	-e SLACK_NOTIFY_TOKEN \
	-e SLACK_NOTIFY_CHANNEL \
	hatappi/slack-notify \
	-attachments '[{"pretext": "pre-hello", "color": "#00CCFF","text": "hello"}]'
```