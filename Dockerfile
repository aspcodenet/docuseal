# Build the Go app
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Create the final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY templates ./templates
EXPOSE 8080
CMD ["./main"]
