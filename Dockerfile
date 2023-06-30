FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . ./
RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/*.proto
RUN CGO_ENABLED=0 GOOS=linux go build cmd/chatedge/main.go

From alpine:latest
WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]
