DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v6.6.0 generate -i sdk/api_files/msc-dcsa-track-and-trace-2.2.yaml
GO_CLIENT := -g go -o /sdk/api \
			--git-repo-id=go-msc-sdk --git-user-id=buyco \
			--additional-properties=packageName=api \
			--additional-properties=isGoSubmodule=false \
			--additional-properties=generateInterfaces=true

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

## test: Run test suites
test: go-get go-test

## check: Format and lint
check: fmt lint go-tidy

## lint: Run go fmt
fmt:
	@echo "  >  Running go formatter..."
	gofmt -w -s ./ 1>&2

## lint: Run linter
lint:
	@echo "  >  Running staticcheck go linter..."
	@GOBIN=$(GOBIN) go install honnef.co/go/tools/cmd/staticcheck@v0.6.1
	@$(GOBIN)/staticcheck ./auth

## lint: Run vet
vet:
	@echo "  >  Running go vet..."
	go vet ./...

go-test:
	@echo "  >  Run tests..."
	@GOBIN=$(GOBIN) go install github.com/onsi/ginkgo/v2/ginkgo@v2.22
	@$(GOBIN)/ginkgo -r --randomize-all --randomize-suites --race --trace -coverprofile=cover.out -gcflags="-l" 1>&2

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	go get ./...

go-tidy:
	@echo "  > Running go mod tidy"
	go mod tidy

## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./api

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
