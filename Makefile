clean:
	@rm -rf ./tmp
	@mkdir ./tmp

packages:
	go mod tidy

build: clean packages lint vet
	mkdir -p out
	go build -o ./out/dago ./cmd/dago/*.go

test:
	go test ./...

vet: 
	@echo "Checking vet..." 
	@go vet ./...

lint: 
	@golangci-lint run
