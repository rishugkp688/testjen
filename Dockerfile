# ----------- STAGE 1: Build Go App -----------
FROM golang:1.24.5 as builder

WORKDIR /app

# Copy backend source
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/*.go ./
RUN go build -o server

# ----------- STAGE 2: Final Image -----------
FROM alpine:latest

# Install SQLite CLI (if needed)
RUN apk --no-cache add sqlite

WORKDIR /root/

# Copy the Go server binary
COPY --from=builder /app/server .

# Copy frontend files
COPY frontend/ ./frontend/

# Expose port
EXPOSE 8080

# Run the server
CMD ["./server"]

