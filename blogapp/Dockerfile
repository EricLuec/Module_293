FROM golang:1.24.2 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./backend

# Ändere hier das Arbeitsverzeichnis
WORKDIR /app/backend

# Baue direkt aus dem Ordner mit main.go
RUN go build -o /app/blogapp

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/blogapp .
COPY frontend/ ./frontend

EXPOSE 8080

CMD ["./blogapp"]
