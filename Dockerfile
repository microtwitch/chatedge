FROM golang:1.20 AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build cmd/chatedge/main.go

From alpine:latest
WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]
