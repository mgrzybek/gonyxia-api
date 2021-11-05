.PHONY: help # This help message
help:
	@grep '^.PHONY: .* #' Makefile \
	| sed 's/\.PHONY: \(.*\) # \(.*\)/\1\t\2/' \
	| expand -t20 \
	| sort

.PHONY: pre-commit # Run pre-commit compliance tests
pre-commit:
	pre-commit run --all-files

.PHONY: test # Run go test
test:
	go test

onyxia-api: test
	go build -o onyxya-api main.go

.PHONY: swagger.validate # Validate the Swagger YAML hile
swagger.validate:
	swagger validate pkg/swagger/swagger.yml

.PHONY: all # lint, test and build
all: onyxia-api
	@echo
