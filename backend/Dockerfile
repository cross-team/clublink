FROM golang:1.13.1-alpine AS builder

WORKDIR /clublink

RUN apk add --no-cache git bash

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY . .

RUN go build -o build/app main.go

FROM alpine:3.10 as production

WORKDIR /clublink

RUN apk add --no-cache bash

COPY --from=builder /clublink/build/app ./build/app
COPY --from=builder /clublink/scripts/wait-for-it ./scripts/wait-for-it
COPY --from=builder /clublink/app/adapter/sqldb/migration ./app/adapter/sqldb/migration
COPY --from=builder /clublink/app/adapter/routing/public ./app/adapter/routing/public
COPY --from=builder /clublink/app/adapter/routing/api.yml ./app/adapter/routing/api.yml
COPY --from=builder /clublink/app/adapter/gqlapi/schema.graphql ./app/adapter/gqlapi/schema.graphql