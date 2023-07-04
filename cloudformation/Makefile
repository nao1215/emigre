ROOT_DIR := $(shell git rev-parse --show-toplevel)

.PHONY: s3
local-s3: ## Deploy the bucket on LocalStack
	awslocal cloudformation deploy --stack-name emigre-local-stack --template-file $(ROOT_DIR)/cloudformation/s3.yml

.DEFAULT_GOAL := help
help:  ## Print help message
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'