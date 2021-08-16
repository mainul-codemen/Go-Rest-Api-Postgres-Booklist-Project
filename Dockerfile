# Specify the base image for the go app.
FROM golang:1.15
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/Go-Rest-Api-Postgres-Booklist-Project    
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq
# Compile the binary exe for our app.
RUN go build -o main .
# Start the application.
CMD ["./main"]

# # Start from golang base image
# FROM golang:alpine as builder

# # ENV GO111MODULE=on

# # Add Maintainer info
# LABEL maintainer="Mainul Hasan <mainul080@gmail.com>"

# # Install git.
# # Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git

# # Set the current working directory inside the container 
# WORKDIR /app

# # Copy go mod and sum files 
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
# RUN go mod download 

# # Copy the source from the current directory to the working Directory inside the container 
# COPY . .

# # Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# # Start a new stage from scratch
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .       

# # Expose port 8080 to the outside world
# EXPOSE 8080

# #Command to run the executable
# #CMD ["./Go-Rest-Api-Postgres-Booklist-Project"]


# CMD ["./main"]