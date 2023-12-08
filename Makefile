.PHONY: run vendor build

# load dotenv & export
-include .env
export

ROOT = $(PWD)

# Install go modules dependencies
vendor:
	go mod vendor
tidy:
	go mod tidy
run:
	@go run ./cmd/scd/main.go

test:
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

up:
	@docker-compose -f docker-compose.local.yaml up --build

down:
	@docker-compose -f docker-compose.local.yaml down -v

build:
	@CGO_ENABLED=0 go build -tags dynamic -o scd ./cmd/scd/main.go

# build for air
build-air:
	@CGO_ENABLED=0 go build -tags dynamic -o ./tmp/app/engine cmd/scd/main.go
