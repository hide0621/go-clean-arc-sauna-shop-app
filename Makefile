.PHONY: help

##################
#### Swagger #####
##################
gen-swagger: ## Generate Swagger documentation with Swaggo
	swag init -g app/cmd/main.go  --output app/docs/swagger

##################
#### help #####
##################
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'