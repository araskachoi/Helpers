FROM golang:1.13-alpine
  
RUN apk update && \
    apk add --no-cache make gcc musl-dev linux-headers git sudo bash tmux
WORKDIR /
COPY . /ethAccGen
WORKDIR ethAccGen
#ENV GOBIN /go/bin
RUN go get -t -v ./...
RUN go build

ENTRYPOINT ["/bin/sh"]
