FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o hello .

#FROM alpine:latest
FROM scratch
WORKDIR /app
COPY --from=builder /app/hello /app/hello
#RUN chmod +x /app/hello
CMD ["./hello"]
