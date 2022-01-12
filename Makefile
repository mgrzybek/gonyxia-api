BINARY      = onyxia-api
LAST_COMMIT = $(shell git rev-parse HEAD)

.PHONY: help
help: ## This help message
	@awk -F: \
		'/^([a-z-]+): [a-z- ]*## (.+)$$/ {gsub(/: .*?\s*##/, "\t");print}' \
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
	go build -o onyxia-api main.go

##############################################################################
# Vagrant

.PHONY: vagrant-variables
vagrant-variables: ## Test vagrant env variables
	@echo -n "Checking VAGRANT_BOX_NAME... "
	@test -z "$$VAGRANT_BOX_NAME" || echo OK

.PHONY: vagrant-destroy
vagrant-destroy: ## Destroy vagrant boxes
	vagrant destroy -f

.PHONY: vagrant-vbox
vagrant-vbox: vagrant-variables ## Test the api using vagrant and virtualbox
	vagrant up --provider=virtualbox
	vagrant provision

##############################################################################
# Containers

.PHONY: oci
oci: onyxia-api ## Create an OCI image using podman build
	podman build --format=oci --tag=${BINARY}:${LAST_COMMIT} .

.PHONY: docker
docker: onyxia-api ## Create a docker image using docker build
	docker build --tag=${BINARY}:${LAST_COMMIT} .

##############################################################################
# All

.PHONY: clean
clean: ## Delete produced artifacts
	rm -f onyxia-api

.PHONY: all
all: get onyxia-api ## Test and build
	@echo
