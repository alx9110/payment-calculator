# Web
FROM node:14-alpine as ui

COPY ./web /app/web

WORKDIR /app/web/payment-calculator

RUN npm install && npm run build

# Backend
FROM golang:1.17-alpine as build

WORKDIR /app

COPY . /app/

RUN apk --no-cache add alpine-sdk
COPY --from=ui /app/web/payment-calculator/dist/routing-app /srv/www
ENV PAYMENT_CALCULATOR_ENV=PRODUCTION
RUN go build
ENTRYPOINT [ "/app/payment-calculator" ]

EXPOSE 8080
