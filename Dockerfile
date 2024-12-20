# Stage 1: Build stage
FROM docker.arvancloud.ir/golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o dgres .

# Stage 2: Final stage
FROM docker.arvancloud.ir/postgres:15-alpine

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/dgres .

## Set the timezone and install CA certificates
#RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/dgres"]