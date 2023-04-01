FROM golang:1.19-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build .

EXPOSE 8080

CMD ["./main"]