.PHONY: help
help: ## This help message
	@awk -F: \
		'/^([a-z-]+): \w*\s*## (.+)$$/ {gsub("[a-z ]*## ","") ; print $$1"\t"$$2}' \
		Makefile \
	| expand -t20 \
	| sort

.PHONY: pre-commit
pre-commit: ## Run pre-commit compliance tests
	pre-commit install
	pre-commit run --all-files

.PHONY: test
test: ## Run go test
	go test

.PHONY: get
get: ## Download required modules
	go get ./...

onyxia-api: test ## Test and build the program
	go build -o onyxya-api main.go

.PHONY: all
all: onyxia-api ## Test and build
	@echo
