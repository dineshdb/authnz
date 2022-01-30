SHELL=/bin/bash
.PHONY: help

list:
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test ./...

vet: ## Vet the code
	go vet ./...

fmt: ## Format the code
	@go mod tidy
	@go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	@test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint: ## Lint the code
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build: ## Build the code
	go build -o bin/server ./cmd/server

run: ## Run the code for development purpose
	go run cmd/server/main.go

start: build ## Build and run the code for production purpose
	bin/server

image: ## Build a docker image
	docker build -t dineshdb/authnz .
	
up: ## Run docker image
	docker run -p 8080:8080 dineshdb/authnz
