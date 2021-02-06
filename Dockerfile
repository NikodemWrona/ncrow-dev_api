FROM golang:1.15-alpine AS builder

RUN apk add --no-cache  git

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN

RUN mkdir -p /go/src/github.com/nikodemwrona/ncrow-dev_api/

RUN go env -w GOPRIVATE="github.com/nikodemwrona/ncrow-dev_api"
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/nikodemwrona".insteadOf "https://github.com/nikodemwrona"

COPY . /go/src/github.com/nikodemwrona/ncrow-dev_api

WORKDIR /go/src/github.com/nikodemwrona/ncrow-dev_api

RUN go mod download

RUN mkdir ./bin/
RUN go build -o ./bin/ ./src/cmd/main.go

FROM alpine

RUN adduser -S -D -H -h /usr/app appuser
USER appuser

COPY --from=builder /go/src/github.com/nikodemwrona/ncrow-dev_api/bin /usr/app/

WORKDIR /usr/app/

CMD ["./main"]
