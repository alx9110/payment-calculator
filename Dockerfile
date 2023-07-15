# Web

COPY ./web /app/web

WORKDIR /app/web/payment-calculator

RUN npm install && npm run build

# Backend
FROM golang:1.17-alpine as build

WORKDIR /app

COPY . /app/

RUN mkdir -p /var/payment-calculator
RUN apk --no-cache add alpine-sdk
COPY --from=ui /app/web/payment-calculator/dist/routing-app /srv/www
ENV PAYMENT_CALCULATOR_ENV=PRODUCTION
RUN go build
ENTRYPOINT [ "/app/payment" ]

EXPOSE 8080
