FROM golang:1.19.2-alpine

RUN apk add --no-cache make

WORKDIR /app

CMD ["make", "dev"]
