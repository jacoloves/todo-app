FROM golang:1.21-alpine AS builder

RUN apk add --no-cache nodejs npm mysql-client

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy

RUN npm install -g prisma
RUN prisma generate

RUN go build -o todo-backend

FROM alpine:latest
RUN apk --no-cache add ca-certificates mysql-client nodejs npm go
WORKDIR /root/
COPY --from=builder /app/todo-backend .
COPY --from=builder /app/wait-for-db.sh .
COPY .env .
COPY prisma /prisma

CMD ["./wait-for-db.sh", "./todo-backend"]
