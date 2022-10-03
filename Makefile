BINDIR			:= $(CURDIR)/bin
INSTALL_PATH	?= /usr/local/bin
BINNAME			?= closeddoors

# go options
PKG        := ./test/...
TAGS       :=
TESTS      := .
TESTFLAGS  := -v
LDFLAGS    := -w -s
GOFLAGS    :=

# Required for globs to work correctly
SHELL      = /usr/bin/env bash

.PHONY: build
build:
	go build -o $(BINDIR)/$(BINNAME) cmd/closeddoors/main.go

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test: build
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	GO111MODULE=on go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)