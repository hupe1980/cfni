PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: setup
## setup: Setup installes dependencies
setup:
	go mod tidy -compat=1.20

.PHONY: test
## test: Runs go test with default values
test:
	go test -v -race -count=1  ./...

.PHONY: build
## build: Builds a version of cfni
build:
	go build -o dist/

.PHONY: ci
## ci: Run all the tests and code checks
ci: build test

.PHONY: cleanup
## cleanup: Runs cfni cleanup
cleanup:
	go run main.go cleanup -b "${BUCKET}"

.PHONY: run
## run: Runs cfni
run:
	go run main.go cfn-code-execution -b "${BUCKET}" -f input.js --runtime nodejs16.x -h

.PHONY: help
## help: Prints this help message
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo