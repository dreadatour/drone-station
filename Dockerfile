# Build binary
# ============
FROM golang:1.11-alpine AS builder

# Install git (required for fetching the dependencies)
RUN apk update && apk add --no-cache git
ADD . /go/src/github.com/dreadatour/drone-station
WORKDIR /go/src/github.com/dreadatour/drone-station

# Build binaries
RUN go install ./...


# Build Docker image
# ==================
FROM alpine:3.8

RUN mkdir /app

# Copy static executables
COPY --from=builder /go/bin/api /app/api

WORKDIR /app

# Run the api binary
ENTRYPOINT ["/app/api"]
