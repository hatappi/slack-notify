FROM golang:1.13-alpine as builder

WORKDIR /go/src/hatappi/slack-notify

RUN apk add make

COPY . .

RUN make build

FROM alpine

COPY --from=builder /go/src/hatappi/slack-notify/dist/slack-notify /usr/local/bin/

RUN adduser slack-notify --disabled-password && \
  adduser slack-notify slack-notify

RUN chown slack-notify:slack-notify /usr/local/bin/slack-notify

USER slack-notify

ENTRYPOINT /usr/local/bin/slack-notify
