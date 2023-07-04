set dotenv-load

clean-proto:
    rm --force protos/*.go

protoc: clean-proto
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/*.proto

run: protoc
    go run cmd/chatedge/main.go

run-example: protoc
    go run cmd/example/main.go

clean:
    rm main

build: protoc
    CGO_ENABLED=0 GOOS=linux go build cmd/chatedge/main.go

debug:
    dlv debug cmd/all/main.go

docker-build: clean-proto
    docker build --tag ghcr.io/microtwitch/chatedge:latest .

docker-release version:
    git checkout {{version}}
    docker build --tag ghcr.io/microtwitch/chatedge:{{version}} .
    docker push ghcr.io/microtwitch/chatedge:{{version}}

lint:
    golangci-lint run
