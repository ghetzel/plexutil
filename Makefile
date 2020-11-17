LOCALS   := $(shell find . -type f -name '*.go')

.PHONY: deps fmt build
.EXPORT_ALL_VARIABLES:

GO111MODULE = on
CGO_ENABLED = 0

all: fmt build

deps:
	go get ./...

fmt:
	go generate -x ./...
	gofmt -w $(LOCALS)
	go vet ./...
	-go mod tidy

build:
	go build -o bin/plexutil cmd/plexutil/*.go
	which plexutil 2> /dev/null && cp -v bin/plexutil $(shell which plexutil 2> /dev/null) || true