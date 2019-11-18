FROM golang:1.13-alpine
  
RUN apk update && \
    apk add --no-cache make gcc musl-dev linux-headers git sudo bash tmux
WORKDIR /
ENV GOBIN /go/bin
RUN git clone https://github.com/araskachoi/Helpers.git
WORKDIR Helpers/ethAccGen
RUN go get -t -v ./...
RUN go build

ENTRYPOINT ["/bin/sh"]
