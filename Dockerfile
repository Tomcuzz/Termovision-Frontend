# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/tomcuzz/Termovision-Frontend/src

# Setup environment veriables
ENV SERVER_ADDRESS=":8080"
ENV BACKEND_ADDRESSES="0.0.0.0"

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src .

# Inatall dependancies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o main main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main", "--server_address ${SERVER_ADDRESS}", "--backend_addresses ${BACKEND_ADDRESSES}"]
