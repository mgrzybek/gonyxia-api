.PHONY: help
help: ## This help message
	@awk -F: \
		'/^([a-z-]+): \w*\s*## (.+)$$/ {gsub("[a-z ]*## ","") ; print $$1"\t"$$2}' \
		Makefile \
	| expand -t20 \
	| sort

##############################################################################
# Tools

.PHONY: pre-commit
pre-commit: ## Run pre-commit compliance tests
	pre-commit install
	pre-commit run --all-files

##############################################################################
# Go

.PHONY: test
test: ## Run go test
	go test

.PHONY: get
get: ## Download required modules
	go get ./...

onyxia-api: test ## Test and build the program
	go build -o onyxya-api main.go

##############################################################################
# Vagrant

.PHONY: vagrant-variables
vagrant-variables: ## Test vagrant env variables
	@echo -n "Checking VAGRANT_BOX_NAME... "
	@[ ! "$$VAGRANT_BOX_NAME" = "" ] && echo OK

.PHONY: vagrant-destroy
vagrant-destroy: ## Destroy vagrant boxes
	@vagrant destroy -f

.PHONY: vagrant-vbox
vagrant-vbox: vagrant-variables ## Test the api using vagrant and virtualbox
	@vagrant up --provider=virtualbox
	@vagrant provision

##############################################################################
# All

.PHONY: all
all: get onyxia-api ## Test and build
	@echo
