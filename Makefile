GOBIN:=${PWD}/bin
PATH:=${GOBIN}:${PATH}

.PHONY: run
run:
	go run cmd/slack-notify/main.go

.PHONY: build
build:
	go build -o dist/slack-notify cmd/slack-notify/main.go

install-tools:
	@GOBIN=${GOBIN} ./scripts/install_tools.sh

lint:
	@${GOBIN}/golangci-lint run ./...
