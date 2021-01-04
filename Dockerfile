FROM  node:14.15-alpine as assets-builder

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY package.json package-lock.json ./
COPY assets/. ./assets/

RUN npm install

# Compile SCSS files
RUN npm run build-scss

# Compile Tailwind CSS files
RUN npx tailwindcss build assets/stylesheets/vendors/tailwind.css -o ./static/css/tailwind.css

# Minify Javascript files
RUN npm run minify-js

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
