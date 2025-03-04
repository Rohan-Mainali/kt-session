FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 9000

# Command to run the executable
CMD ["/app/main"]

