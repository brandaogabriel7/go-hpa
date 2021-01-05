FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-hpa

FROM scratch

COPY --from=builder /go/bin/go-hpa /go/bin/go-hpa

ENTRYPOINT ["/go/bin/go-hpa"]

EXPOSE 8080