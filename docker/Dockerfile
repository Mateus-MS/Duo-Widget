FROM golang:1.25.4 AS builder

WORKDIR /app

# Copy all files into container
COPY ../pkg .

# Compile go
RUN go mod tidy
RUN go build -o server .

# Image that gonna run
FROM debian:bookworm-slim
WORKDIR /app

# Install CA certificates so Go can verify HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Move the compiled go to the run-time image
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"] 