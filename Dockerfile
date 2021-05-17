FROM  node:14.15-alpine AS assets-builder

RUN apk --no-cache add ca-certificates make

WORKDIR /app

COPY package.json package-lock.json ./
COPY assets/. ./assets/

ADD .env.example ./.env
COPY Makefile ./Makefile

RUN npm install

# Prepare all assets
RUN make assets

FROM golang:1.15-alpine AS migration

ARG DATABASE_URL

# Move to working directory /migration
WORKDIR /migration

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy the code into the container
COPY . .

# Install command-line tool
RUN go get github.com/beego/bee/v2

# Migrate database
RUN bee migrate -driver=postgres -conn=$DATABASE_URL

FROM golang:1.15-alpine

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies
RUN go mod download

# Copy ENV file
ADD .env.example ./.env

# Copy the code into the container
COPY . .

# Copy assets from assets builder
COPY --from=assets-builder /app/static/. ./static/

# Build the application
RUN go build -o main .

EXPOSE 8080

# Run the executable
CMD ["./main"]
