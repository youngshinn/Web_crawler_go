# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 
RUN go build -o crawler-app .

# Stage 2: Runtime
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/crawler-app /app/crawler-app
RUN chmod +x /app/crawler-app

ENTRYPOINT ["/app/crawler-app"]
CMD ["Golang", "2"]
