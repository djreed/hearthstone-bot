FROM golang:1.8

# Add the source code:
WORKDIR /app/
ADD oauth/ ./oauth
ADD battlenet/ ./battlenet
ADD slack/ ./slack
ADD sanitize/ ./sanitize
ADD certs/ ./certs
ADD ./*.go ./

# Build it:
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GODEBUG=netdns=cgo go build -a -tags netgo -ldflags '-w' -o hearthstone-bot .

# Executable container
FROM alpine

WORKDIR /app/

COPY --from=0 /app/hearthstone-bot .

CMD ["/app/hearthstone-bot"]
