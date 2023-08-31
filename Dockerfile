# Use an official Golang runtime as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Golang application
RUN go build -o myapp

# Specify the command to run your application
CMD ["./myapp"]
