FROM golang:1.20.5-alpine3.18 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go get
RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]