FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base git inotify-tools
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN echo $GOPATH

COPY . /clean_web
WORKDIR /clean_web

RUN go mod download

CMD sh /clean_web/docker/run.sh
