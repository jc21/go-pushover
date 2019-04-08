# Go parameters
VERSION=$(shell cat package.json | jq -r .version)
GITCOMMIT=$(shell git rev-list -1 HEAD)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=pushover

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags "-X main.GitCommit=$(GITCOMMIT) -X main.AppVersion=$(VERSION)"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps:
	$(GOGET)
