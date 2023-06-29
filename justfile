set dotenv-load

clean-proto:
    rm protos/*.go

protoc: clean-proto
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/*.proto

run: protoc
    go run cmd/chatedge/main.go

run-example: protoc
    go run cmd/example/main.go

debug:
    dlv debug cmd/chatedge/main.go

docker-build: protoc
    docker build --tag ghcr.io/microtwitch/chatedge:latest .

docker-push: docker-build
    docker push ghcr.io/microtwitch/chatedge:latest
