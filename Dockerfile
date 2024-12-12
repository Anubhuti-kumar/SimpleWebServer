# Use an official Go image
FROM golang:1.20-alpine

# Set the current working directory
WORKDIR /app

# Copy all files to the container
COPY . .

# Install dependencies and build the application
RUN go mod tidy

# Command to run your application
CMD ["./main"]

