# Gunakan image Go untuk build
FROM golang:1.22-alpine as builder

# Set working directory di dalam container
WORKDIR /app

# Copy go mod dan go sum file
COPY go.mod go.sum ./

# Download dependensi yang dibutuhkan
RUN go mod tidy

# Copy seluruh file sumber aplikasi ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Aplikasi siap dijalankan menggunakan image yang lebih ringan
FROM alpine:latest  

# Install PostgreSQL client (opsional)
RUN apk --no-cache add postgresql-client

# Set working directory di dalam container
WORKDIR /root/

# Copy file binary dari builder image
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 


# Expose port yang digunakan aplikasi Go
EXPOSE 3000

# Perintah untuk menjalankan aplikasi
CMD ["./main"]
