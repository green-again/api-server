default: build

help:
	@echo 'Management commands for api-server:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make up              Start the service.'
	@echo '    make tidy            Prune any extraneous requirements.'
	@echo '    make lint            Lint project.'
	@echo '    make test            Run tests on a compiled project.'
	@echo

build:
	go build ./cmd/server.go

up:
	go run ./cmd/server.go

tidy:
	go mod tidy

lint:
	go vet ./...

test:
	go test ./...
