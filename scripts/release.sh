#!/bin/bash -e

REPO=hatappi/slack-notify

if [ -z "$GITHUB_TOKEN" ]; then
  echo "Please set GITHUB_TOKEN"
  exit 1
fi

rsync -avr --exclude='.envrc' --exclude='tmp' . ./tmp

docker run --rm \
  -v $PWD/tmp:/go/src/github.com/${REPO} \
  -w /go/src/github.com/${REPO} \
  -e GITHUB_TOKEN \
  -it alpine sh

#goreleaser/goreleaser release
