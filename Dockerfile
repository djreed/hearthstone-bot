FROM golang:1.12

# Add the source code:
WORKDIR /go/src/github.com/djreed/hearthstone-bot/
ADD oauth ./oauth
ADD battlenet/ ./battlenet
ADD slack/ ./slack
ADD sanitize/ ./sanitize
ADD ssl/ ./ssl
ADD ./*.go go.mod go.sum ./

# Build it:
RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o hearthstone-bot .

# Executable container
FROM alpine

WORKDIR /app/

COPY --from=0 /go/src/github.com/djreed/hearthstone-bot/hearthstone-bot .

CMD ["/app/hearthstone-bot"]
