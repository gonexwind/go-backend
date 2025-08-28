# Base image
FROM golang:1.25-alpine

# Set working directory
WORKDIR /app

# Copy go files
COPY . .

# Build Go app
RUN go mod init gonexwind/backend-api && go build -o main .

# Run app
CMD ["./main"]

EXPOSE 8080
