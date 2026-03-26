# syntax=docker/dockerfile:1.7

FROM golang:1.23-bookworm AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/shiro-bot ./cmd/bot
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/deploy-commands ./cmd/cli/deploy-commands

FROM gcr.io/distroless/base-debian12:nonroot AS runtime
WORKDIR /app

COPY --from=builder /out/shiro-bot /app/shiro-bot
COPY --from=builder /out/deploy-commands /app/deploy-commands

ENV ENVIRONMENT=production

ENTRYPOINT ["/app/shiro-bot"]
