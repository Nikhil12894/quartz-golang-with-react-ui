# Build executable stage
FROM golang
# Add Maintainer Info
LABEL maintainer="Nalin <nalin12894@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY /server/go.mod /server/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY /server/*.go .
COPY /server/public ./public

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Build final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=0 /app .

# Command to run
ENTRYPOINT ["/main"]