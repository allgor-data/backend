FROM golang:1.19.2-alpine

RUN apk add --no-cache make

RUN addgroup -S nonroot \
  && adduser -S nonroot -G nonroot

WORKDIR /app

USER nonroot

CMD ["make", "dev"]
