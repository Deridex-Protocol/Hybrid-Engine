FROM golang:1.17-alpine as builder

RUN apk add build-base git

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY . .

# CGO_ENABLED=0 for scratch
RUN CGO_ENABLED=0 go build ./cmd/bot/main.go

# multi stage build
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/main /app/

CMD ["/app/main"]