# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .

EXPOSE 8080

# Run
CMD ["go","run","main.go"]