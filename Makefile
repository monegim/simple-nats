APP_NAME = simple-nats
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
build:
	docker build -t ${APP_NAME} .
build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

up: build
	docker compose up -d

down:
	docker compose down
