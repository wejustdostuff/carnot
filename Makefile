default: build

PKGS=$(shell go list ./... | grep -v /vendor/)
GORELEASER := $(shell command -v goreleaser 2> /dev/null)

.PHONY: dep vet style format test build release goreleaser

all: dep vet style format test build

dep:
	go mod download

vet:
	@go vet $(PKGS)

style:
	@gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@go fmt $(PKGS)

test:
	go test -v -cover -race $(PKGS)

build:
ifndef GORELEASER
	$(error "goreleaser not found (`go get -u -v github.com/goreleaser/goreleaser` to fix)")
endif
	$(GORELEASER) --skip-publish --rm-dist --snapshot

release:
ifndef GORELEASER
	$(error "goreleaser not found (`go get -u -v github.com/goreleaser/goreleaser` to fix)")
endif
	$(GORELEASER) --rm-dist

goreleaser:
	go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser
