FROM golang:1.25-alpine

ENV TZ=Asia/Tokyo
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download


COPY . .

EXPOSE ${GO_PORT}

CMD ["air"]