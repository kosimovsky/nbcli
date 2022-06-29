# Makefile

BIN ?= nbcli

local : ARCH ?= $(shell go env GOOS)-$(shell go env GOARCH)
ARCH ?= linux-amd64

platform_temp = $(subst -, ,$(ARCH))
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))

all:
	@$(MAKE) local

local: build-dirs
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	BIN=$(BIN) \
	OUTPUT_DIR=$$(pwd)/bin/$(GOOS)/$(GOARCH) \
	./build/build.sh

build: bin/$(GOOS)/$(GOARCH)/$(BIN)

bin/$(GOOS)/$(GOARCH)/$(BIN): build-dirs
	@echo "building: $@"
	$(MAKE) shell CMD="-c '\
		GOOS=$(GOOS) \
		GOARCH=$(GOARCH) \
		BIN=$(BIN) \
		OUTPUT_DIR=$$(pwd)/bin/$(GOOS)/$(GOARCH) \
		./build/build.sh'"

build-dirs:
	@mkdir -p bin/$(GOOS)/$(GOARCH)

shell:
	/bin/sh $(CMD)

clean:
	rm -rf  bin


.PHONY: modules
modules:
	go mod tidy