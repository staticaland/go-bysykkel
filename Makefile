include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run/cli: run the cmd/bysykkel application
.PHONY: run/cli
run/cli:
	go run ./cmd/bysykkel

## run/server: run the cmd/bysykkel-serve application
.PHONY: run/server
run/server:
	go run ./cmd/bysykkel-serve

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
	go test -race -vet=off ./...
