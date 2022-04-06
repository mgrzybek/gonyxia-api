# gonyxia-api
A rewrite of https://github.com/InseeFrLab/onyxia-api using Go Language


## Build and Run

The tasks are available through a Makfile.

```bash
$ make help
all                  Test and build
clean                Delete produced artifacts
docker               Create a docker image using docker build
get                  Download required modules
help                 This help message
oci                  Create an OCI image using podman build
onyxia-api           Test and build the program
pre-commit           Run pre-commit compliance tests
run-in-cluster       Run the server on 127.0.0.1:8081 using in-cluster
run-out-cluster      Run the server on 127.0.0.1:8081 using out-cluster
test                 Run go test
vagrant-destroy      Destroy vagrant boxes
vagrant-variables    Test vagrant env variables
vagrant-vbox         Test the api using vagrant and virtualbox
```

### Create a developer environment

```bash
# Install go and pre-commit
$ make get pre-commit
# Test the code
$ make test
```

### Start a local instance

```bash
# Start a local instance in in-cluster mode
$ make run-in-cluster
# Start a local instance using a remote orchestrator
$ make run-out-cluster
```

### Using Vagrant

A k3s Vagrant box can be created in order to user the *out-cluster* target.
However this functionnality is a WIP.

```bash
# Start the box using Virtualbox
$ make vagrant-box
# TODO: Set credentials in the region's configuration
# Start a local instance using a remote orchestrator
$ make run-out-cluster
```


## Create a configuration file

For each region, some ID provider attributes has to be provided. Some attributes can be read from the environment.
* `AUTH_REALM`
* `AUTH_SERVER_URL`
* `AUTH_REDIRECT_URL`
* `AUTH_CLIENT_ID`
* `AUTH_CLIENT_SECRET`
