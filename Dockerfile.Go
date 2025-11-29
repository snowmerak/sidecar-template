# ==========================================
# Dockerfile.Go
# ==========================================
FROM golang:1-alpine AS builder

# 1. Install build dependencies
# git is needed for go mod download if dependencies are private or git-based
RUN apk add --no-cache git

WORKDIR /app

# 2. Copy source code (including gen/ folder)
COPY . .

# 3. Build Sidecar Go
WORKDIR /app/sidecar-go
RUN go mod tidy
RUN go build -o /app/bin/sidecar-go main.go

# ==========================================
# Runner
# ==========================================
FROM alpine:3

WORKDIR /app

COPY --from=builder /app/bin/sidecar-go /app/sidecar-go

EXPOSE 50051

CMD ["/app/sidecar-go"]
