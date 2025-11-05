FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:latest

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/main .
EXPOSE ${GO_PORT}
CMD ["./main"]