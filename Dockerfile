# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . 

# Build the golang-docker command inside the container.
RUN go install  

# Run the golang-docker command by default when the container starts.
ENTRYPOINT  

# http server listens on port 8080.
EXPOSE 8080