FROM golang:1.13.1-alpine AS builder

WORKDIR /clublink

RUN apk add --no-cache git bash

RUN git config \
  --global \
  url."https://MarcelloPaiva:dc87ee1de6a370db146cb294518b2f1e772052f4@github.com/".insteadOf \
  "https://github.com/"

# Install dependencies
COPY serve/go.mod serve/go.sum ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY serve .

RUN go build -o app main.go

FROM alpine:3.10 as production

WORKDIR /clublink

RUN apk add --no-cache bash

COPY ./build ./public

COPY --from=builder /clublink/app ./app

ENTRYPOINT ./app