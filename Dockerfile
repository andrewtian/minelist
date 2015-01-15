FROM golang:1.4

ADD . /go/src/github.com/andrewtian/minelist/

WORKDIR /go/src/github.com/andrewtian/minelist/

RUN go get
RUN go build