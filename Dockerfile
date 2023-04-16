# Use an official Golang runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Install any needed dependencies specified in go.mod
RUN go mod download

# Build the app
RUN go build -o app cmd/api/main.go

# Expose port 8080 for the API
EXPOSE 8080

# Run the app when the container starts
CMD ["./app"]
