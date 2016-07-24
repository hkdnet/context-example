FROM golang:1.6.3-alpine
ADD . $GOPATH/src/github.com/hkdnet/context-example

CMD ["go", "run", "/go/src/github.com/hkdnet/context-example/main.go"]


