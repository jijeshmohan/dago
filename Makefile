.POSIX:
SHELL := /bin/bash # Use bash syntax

.EXPORT_ALL_VARIABLES:

BIN_NAME=dago

clean:
	@rm -rf ./tmp
	@rm -rf ./out
	@mkdir ./tmp

packages:
	go mod tidy

build: clean packages lint vet
	mkdir -p out
	go build -o ./out/$(BIN_NAME) ./cmd/$(BIN_NAME)/*.go

test:
	go test --cover ./...

vet: 
	@echo "Checking vet..." 
	@go vet ./...

lint: 
	@golangci-lint run

build-all: clean test
	GOOS=windows GOARCH=amd64 go build -o ./out/dago_windows_amd64.exe ./cmd/dago/*.go
	GOOS=windows GOARCH=386 go build -o ./out/dago_windows_386.exe ./cmd/dago/*.go
	GOOS=darwin GOARCH=amd64 go build -o ./out/dago_darwin_amd64 ./cmd/dago/*.go
	GOOS=darwin GOARCH=arm64 go build -o ./out/dago_darwin_arm64 ./cmd/dago/*.go
	GOOS=linux GOARCH=amd64 go build -o ./out/dago_linux_amd64 ./cmd/dago/*.go
	GOOS=linux GOARCH=386 go build -o ./out/dago_linux_386 ./cmd/dago/*.go

release: build-all
	$(shell ./scripts/zip_output.sh)