BINARY      = onyxia-api
LAST_COMMIT = $(shell git rev-parse HEAD)

.PHONY: help
help: ## This help message
	@awk -F: \
		'/^([a-z-]+): [a-z/- ]*## (.+)$$/ {gsub(/: .*?\s*##/, "\t");print}' \
		Makefile \
	| expand -t20 \
	| sort

##############################################################################
# Tools

/usr/local/go/bin/go: ## Install go environment (not included in the help message)
	which go || ( \
	sudo apt update \
    && sudo apt -y install pre-commit golint \
    && \
	wget https://golang.org/dl/go1.17.linux-amd64.tar.gz \
    && sudo tar -zxvf go1.17.linux-amd64.tar.gz -C /usr/local/ \
    && rm -f go1.17.linux-amd64.tar.gz \
    && echo "export PATH=/usr/local/go/bin:${PATH}" | sudo tee /etc/profile.d/go.sh \
    && echo "export PATH=/usr/local/go/bin:${PATH}" | sudo tee -a ${HOME}/.profile \
    && echo "\nPlease reload your .profile file to get an updated PATH" )

.PHONY: pre-commit
pre-commit: ## Run pre-commit compliance tests
	pre-commit install
	pre-commit run --all-files

##############################################################################
# Go

.PHONY: test
test: ## Run go test
	go test ./...

.PHONY: get
get: /usr/local/go/bin/go ## Download required modules
	go get ./...

onyxia-api: test ## Test and build the program
	go build -o ${BINARY} main.go

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
# Run (local)

.PHONY: run-in-cluster
run-in-cluster: onyxia-api ## Run the server on 127.0.0.1:8081 using in-cluster
	./${BINARY} \
		-l trace \
		server \
			-b 127.0.0.1:8081 \
			-r ./etc/regions.in-cluster.json \
			-c etc/catalogs.json

.PHONY: run-out-cluster
run-out-cluster: onyxia-api ## Run the server on 127.0.0.1:8081 using out-cluster
	./${BINARY} \
		-l trace \
		server \
			-b 127.0.0.1:8081 \
			-r ./etc/regions.out-cluster.json \
			-c etc/catalogs.json

##############################################################################
# All

.PHONY: clean
clean: ## Delete produced artifacts
	rm -f ${BINARY}

.PHONY: all
all: get onyxia-api ## Test and build
	@echo
