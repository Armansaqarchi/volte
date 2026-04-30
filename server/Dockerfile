# ---- Build stage ----
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Install git (sometimes needed for private/public modules)
RUN apk add --no-cache git ca-certificates

ENV GOPROXY=direct
# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .

# Build a small static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o server .

# ---- Run stage ----
FROM alpine:3.19

WORKDIR /app

# TLS certs for HTTPS calls
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

# Run
ENTRYPOINT ["run.sh"]