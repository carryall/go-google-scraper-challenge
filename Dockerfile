FROM golang:1.18-alpine

ENV GIN_MODE=release

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy codebase
COPY . .

# Build go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

EXPOSE 8080

CMD ["./bin/start.sh"]
