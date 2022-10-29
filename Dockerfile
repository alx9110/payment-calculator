# API
FROM golang:1.17-alpine as build

WORKDIR /app

COPY . /app/

RUN apk --no-cache add alpine-sdk
RUN go build
ENTRYPOINT [ "/app/payment-calculator" ]

EXPOSE 8080
