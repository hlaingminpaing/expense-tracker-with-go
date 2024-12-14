# Use a lightweight base image with Go installed
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the container's /app directory
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Start the application
CMD ["./main"]
