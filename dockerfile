FROM golang:1.17-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build

FROM alpine:3.13.6

COPY --from=builder ["/build/go-healthcheck", "/"]
COPY --from=builder ["/build/test.csv", "/"]

ENTRYPOINT ["/go-healthcheck"]
