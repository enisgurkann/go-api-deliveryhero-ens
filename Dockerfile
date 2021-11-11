# Build go-api-deliveryher-ens in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /go-api
RUN cd /go-api && go mod download && go build

# Pull go-api into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-api/go-api /usr/local/bin/

EXPOSE 8888
ENTRYPOINT ["go-api"]