FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o todo-backend

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/todo-backend .
COPY .env .

CMD ["./todo-backend"]
