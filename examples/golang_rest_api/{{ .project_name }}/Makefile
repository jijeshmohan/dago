BIN_NAME={{.project_name}}

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