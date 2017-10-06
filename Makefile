ifeq ($(OS),Windows_NT)
    ARCH = windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
			ARCH = linux
    endif
    ifeq ($(UNAME_S),Darwin)
			ARCH = darwin
    endif
endif

REPO_VERSION := $$(git describe --abbrev=0 --tags)
BUILD_DATE := $$(date +%Y-%m-%d-%H:%M)
GIT_HASH := $$(git rev-parse --short HEAD)
GOBUILD_VERSION_ARGS := -ldflags "-s -X main.Version=$(REPO_VERSION) -X main.GitCommit=$(GIT_HASH) -X main.BuildDate=$(BUILD_DATE)"
BINARY := neb
MAIN_PKG := github.com/fdesjardins/neb

dep:
	dep ensure

build:
	go build -i -v -o build/bin/$(ARCH)/$(BINARY) $(GOBUILD_VERSION_ARGS) $(MAIN_PKG)

run: build
	./build/bin/$(ARCH)/$(BINARY)

install:
	go install $(GOBUILD_VERSION_ARGS) $(MAIN_PKG)

clean:
	rm -rf build/bin/*

.PHONY: build install
