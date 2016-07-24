FROM golang:1.6.3-alpine
ADD ./context $GOPATH/src/golang.org/x/net/context
ADD . $GOPATH/src/github.com/hkdnet/context-example

CMD ["go", "run", "/go/src/github.com/hkdnet/context-example/main.go"]


