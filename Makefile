MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help
.ONESHELL:


.PHONY: test ## Run the tests
test:
	go test ./...

.PHONY: vtest  ## Run the tests in verbose mode
vtest:
	go test -v ./...

.PHONY:
help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.* ##/:/g' | column -t -s:
