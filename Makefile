# The binary to build (just the basename).
GOFILES      ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*/mock/*")

all: mockgen build run build-tool lint

.PHONY: build
build: ## Build application
	go build main.go

.PHONY: run
run: ## Run application server
	go run main.go

.PHONY: build-tool
build-tool: ## Build application tools
	go build -o cmd/jellyfish_tool/create_user cmd/jellyfish_tool/create_user.go

.PHONY: lint
lint: ## Lint the files
	@echo run lint...
	golangci-lint run

.PHONY: test
test:  ## Run unittests and data race detector
	go test -race -short ./...

.PHONY: mockgen
mockgen: ## generate interfaces mock
	mockgen -source domain/user/repository/user_repository.go -destination domain/user/repository/mock/user_repository.go -package mock
	mockgen -source domain/visitor/repository/visitor_repository.go -destination domain/visitor/repository/mock/visitor_repository.go -package mock
	mockgen -source domain/taco/repository/taco_repository.go -destination domain/taco/repository/mock/taco_repository.go -package mock

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################
.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

log-%:
	@ grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m==> %s\033[0m\n", $$2}'
