.PHONY: run
run:
	go run cmd/slack-notify/main.go

.PHONY: build
build:
	go build -o dist/slack-notify cmd/slack-notify/main.go
