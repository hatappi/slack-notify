#!/bin/bash -e

REPO=hatappi/slack-notify

if [ -z "$GITHUB_TOKEN" ]; then
  echo "Please set GITHUB_TOKEN"
  exit 1
fi

rsync -avr --exclude='.envrc' --exclude='tmp' . ./tmp

docker run --rm \
  -v $PWD/tmp:/go/src/github.com/${REPO} \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /go/src/github.com/${REPO} \
  -e DOCKER_USERNAME \
  -e DOCKER_PASSWORD \
  -e GITHUB_TOKEN \
  goreleaser/goreleaser release --rm-dist
