.PHONY: build test clean

APP         = emigre
VERSION     = $(shell git describe --tags --abbrev=0)
GIT_REVISION := $(shell git rev-parse HEAD)
GO          = go
GO_BUILD    = $(GO) build
GO_TEST     = $(GO) test -v
GO_TOOL     = $(GO) tool
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = -ldflags '-X github.com/nao1215/emigre/server/version.Version=${VERSION}' -ldflags "-X github.com/nao1215/emigre/server/version.Revision=$(GIT_REVISION)"

build:  ## Build binary
	cd server && env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(GO_LDFLAGS) -o $(APP) main.go
	mv server/$(APP) .

run-server: ## Run server
	cd server && env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run main.go

generate: ## Generate file automatically
	docker-compose up -d db
	$(GO) generate ./...
	cd server && swag init 
	sqlc generate --file server/app/schema/sqlc.yml 
	tbls doc --force 
	cd server && wire ./...

clean: ## Clean project
	-rm -rf $(APP) cover.out cover.html
	cd client && gradle clean

test-server: ## Start unit test for server
	cd server && env GOOS=$(GOOS) $(GO_TEST) -cover $(GO_PKGROOT) -coverprofile=coverage.out
	cd server && $(GO_TOOL) cover -html=coverage.out -o coverage.html
	mv server/coverage.* .

test-client: ## Start unit test for client
	cd client && gradle test

create-local-s3: ## Create local s3
	docker-compose up -d localstack
	$(MAKE) -f cloudformation/Makefile local-s3

run-client: ## Run client
	cd client && gradle run

.DEFAULT_GOAL := help
help:  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'