set dotenv-load

run:
    go run cmd/chatedge/main.go

run-example:
    go run cmd/example/main.go

debug:
    dlv debug cmd/chatedge/main.go

docker-build:
    docker build --tag ghcr.io/microtwitch/chatedge:latest .

docker-push: docker-build
    docker push ghcr.io/m4tthewde/chatedge:latest
