# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# create a working directory
WORKDIR /app
# add source code
COPY . .

RUN go build -o main main.go

# run main.go
CMD ["/app/main"]


 
