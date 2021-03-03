# FROM golang:1.13.1-alpine AS builder

# WORKDIR /clublink

# RUN apk add --no-cache git bash

# RUN git config \
#   --global \
#   url."https://MarcelloPaiva:dc87ee1de6a370db146cb294518b2f1e772052f4@github.com/".insteadOf \
#   "https://github.com/"

# # Install dependencies
# COPY serve/go.mod serve/go.sum ./
# RUN go mod download

# # Verify dependencies
# RUN go mod verify

# COPY serve .

# RUN go build -o app main.go

FROM node:14.16.0-alpine3.10 AS npmBuilder

WORKDIR /clublink

COPY . .

RUN yarn

RUN yarn build:production

# FROM alpine:3.10 as production

# WORKDIR /clublink

# RUN apk add --no-cache bash

# COPY --from=npmBuilder /clublink/build ./public

# COPY --from=builder /clublink/app ./app

# ENTRYPOINT ./app

FROM nginx:1.17.5-alpine

WORKDIR /usr/share/nginx/html

COPY --from=npmBuilder /clublink/build .

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY default.conf /etc/nginx/conf.d/