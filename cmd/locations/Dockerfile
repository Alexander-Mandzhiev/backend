FROM golang:alpine AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/locations/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
RUN chmod +x ./main
ENV CONFIG_PATH=config/locations/production.yaml
CMD ["./main"]