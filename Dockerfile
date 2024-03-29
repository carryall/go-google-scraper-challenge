FROM golang:1.18-alpine

ENV GIN_MODE=release

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy codebase
COPY . .

# Copy ENV file
ADD .env.example ./.env

# Build go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

EXPOSE 8080

RUN chmod +x ./bin/start.sh
CMD ["./bin/start.sh"]
