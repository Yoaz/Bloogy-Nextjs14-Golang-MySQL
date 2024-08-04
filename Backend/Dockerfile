# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# # Build the application
# RUN go build -o bin/main 

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

# Expose the port the application runs on
EXPOSE 4000

# Run the application
CMD ["/bin/main"]
