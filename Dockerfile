# Build the Go binary
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/cmd/gobot/gobot ./cmd/gobot

# Set the working directory to where the executable and JSON files are located
WORKDIR /app/cmd/gobot

# Command to run the executable
CMD ["./gobot"]