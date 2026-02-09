# -------------------------
# Build stage
# -------------------------
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata build-base

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o lms-api .

# -------------------------
# Runtime stage
# -------------------------
FROM gcr.io/distroless/base-debian12

ENV TZ=Asia/Kolkata

WORKDIR /app

COPY --from=builder /app/lms-api /app/lms-api

EXPOSE 4008

USER nonroot:nonroot

ENTRYPOINT ["/app/lms-api"]
