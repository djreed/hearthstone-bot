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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hearthstone-bot .

# Executable container
FROM alpine

WORKDIR /app/

COPY --from=0 /app/ .

CMD ["/app/hearthstone-bot"]
