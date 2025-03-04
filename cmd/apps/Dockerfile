FROM golang:alpine AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/apps/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
RUN chmod +x ./main
ENV CONFIG_PATH=config/apps/production.yaml
CMD ["./main"]