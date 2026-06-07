# syntax=docker/dockerfile:1.7

FROM node:22.20-alpine3.21 AS assets

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY tailwind.config.js ./
COPY assets/css/input.css ./assets/css/input.css
COPY assets/js ./assets/js
COPY internal ./internal
RUN npm run css

FROM golang:1.26-alpine3.22 AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020

COPY . .
COPY --from=assets /app/assets/css/app.css ./assets/css/app.css

RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/server ./cmd

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata \
  && addgroup -S app \
  && adduser -S -G app app

COPY --from=builder /out/server ./server
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/static ./static
COPY --from=builder /app/internal/posts ./internal/posts

ENV ADDR=:8080
ENV BASE_URL=https://iphantasmic.com

EXPOSE 8080

USER app

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8080/healthz || exit 1

CMD ["./server"]
