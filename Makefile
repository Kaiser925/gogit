GO:=go

.PHONY: all
all: build

## build: Build code for host platform
.PHONY: build
build:
	@$(GO) build

## test: Run unit test
.PHONY: test
test:
	@${GO} test -v ./...

## install: Install gogit to system
.PHONY: install
install:
	@${GO} install

## remove: Remove gogit from system
.PHONY: remove
remove:
	@-rm ${GOPATH}/bin/gogit

## help: Show help messages
.PHONY: help
help: Makefile
	@echo "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## format: Format all code by using go fmt and golines
.PHONY: format
format: tools.verify.golines
	@${GO} fmt ./...
	@golines -w .

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then ${MAKE} install.$*; fi

.PHONY: install.golines
install.golines:
	@${GO} get -u github.com/segmentio/golines