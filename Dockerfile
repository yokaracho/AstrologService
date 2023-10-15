FROM golang:1.21.2 AS builder

WORKDIR /apod

COPY . .
RUN go build -o astrologservice


FROM alpine:3.14

WORKDIR /apod

COPY --from=builder /apod/astrologservice .

RUN chmod +x /apod/astrologservice

RUN apk add --no-cache ca-certificates

CMD ["./apod/astrologservice"]