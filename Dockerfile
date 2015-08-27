FROM golang:1.5

ADD . $GOPATH/src/github.com/netbrain/gotcptest

RUN go install github.com/netbrain/gotcptest
