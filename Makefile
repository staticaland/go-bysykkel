## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run/cli: run the cmd/bysykkel CLI application
.PHONY: run/cli
run/cli:
	go run ./cmd/bysykkel

## run/server: run the cmd/bysykkel-server REST API application
.PHONY: run/server
run/server:
	go run ./cmd/bysykkel-server

## build/cli: build the cmd/bysykkel CLI application
.PHONY: build/cli
build/cli:
	@echo 'Building cmd/bysykkel...'
	go build -ldflags='-s' -o=./bin/bysykkel ./cmd/bysykkel

## build/server: build the cmd/bysykkel-server REST API application
.PHONY: build/server
build/server:
	@echo 'Building cmd/bysykkel-server...'
	go build -ldflags='-s' -o=./bin/bysykkel-server ./cmd/bysykkel-server

## quality: run quality checks
.PHONY: quality
quality:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -v -race -vet=off ./...
