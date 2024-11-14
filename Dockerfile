FROM golang:1.21.0-alpine AS builder

RUN apk update && apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o go-api main.go

# Run Stage
FROM alpine:latest

COPY --from=builder /app/go-api /go-api

COPY --from=builder /app/templates /templates

RUN chmod +x /go-api

EXPOSE 8080

ENTRYPOINT ["/go-api"]

# CMD ["./go-api"]