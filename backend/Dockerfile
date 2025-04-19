# Gunakan image Go
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy semua file
COPY . .

# Build aplikasi
RUN go build -o main .

# Jalankan
CMD ["./main"]
