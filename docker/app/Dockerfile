FROM golang:1.16-alpine

ENV GOPATH=/go \
    GO111MODULE=on

WORKDIR $GOPATH/src/github.com/nagaihiroya/lets-golang-gin-docker

COPY . $GOPATH/src/github.com/nagaihiroya/lets-golang-gin-docker

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
