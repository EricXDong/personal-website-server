FROM golang:1.10.7-alpine3.7 as development

# Install git
RUN apk update
RUN apk add git

ENV GOPATH /go:/go/src/personal-website-server

RUN mkdir -p /go/src/personal-website-server
WORKDIR /go/src/personal-website-server

# `...` tells go to install all its dependencies too
RUN go get -u github.com/golang/dep/...
RUN go get -v github.com/canthefason/go-watcher/cmd/watcher

ADD . .
RUN dep ensure -v

EXPOSE 5000
RUN go install

ENTRYPOINT [ "watcher" ]

# PRODUCTION BUILD

FROM golang:1.10.7-alpine3.7

RUN mkdir /app
WORKDIR /app

ADD ./config ./config
COPY --from=development /go/bin/personal-website-server .

ENTRYPOINT [ "./personal-website-server" ]
