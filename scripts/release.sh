#!/bin/bash -e

REPO=hatappi/slack-notify

if [ -z "$GITHUB_TOKEN" ]; then
  echo "Please set GITHUB_TOKEN"
  exit 1
fi

docker run --rm \
  -v $PWD:/go/src/github.com/${REPO} \
  -w /go/src/github.com/${REPO} \
  -e GITHUB_TOKEN \
  goreleaser/goreleaser release