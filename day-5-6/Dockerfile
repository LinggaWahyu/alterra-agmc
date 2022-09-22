FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /go-api

FROM alpine:3.16.0

WORKDIR /

COPY --from=builder /go-api .

COPY .env* .

EXPOSE 8080

ENTRYPOINT ["./go-api"]