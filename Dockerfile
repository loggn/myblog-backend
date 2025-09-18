FROM golang:1.25 AS builder


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=builder /app/server .

ENV TZ=Asia/Shanghai

CMD ["./server"]