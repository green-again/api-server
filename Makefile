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
	docker-compose build

up:
	docker-compose up

run:
	docker-compose run --rm app $(cmd)

tidy: cmd=go mod tidy
tidy: run

test: cmd=go test -cover ./...
test: run

test-native:
	go test -cover ./...

lint: cmd=go vet ./...
lint: run

migrate: cmd=go run ./cmd/migrations/main.go
migrate: run

swag:
	swag init -g ./cmd/server/main.go
