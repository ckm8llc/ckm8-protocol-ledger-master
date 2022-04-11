FROM golang:1.11

VOLUME ["/go"]

WORKDIR /go/src/github.com/ckm8token/ckm8/

ENV GOPATH=/go

ENV CGO_ENABLED=1 

ENV GO111MODULE=on

CMD ["/go/src/github.com/ckm8token/ckm8/integration/docker/build/start.sh"]



