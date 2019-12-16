.PHONY: build clean test test-race

ifneq ("$(wildcard /go)", "")
	GOPATH=/go
endif

GOROOT=$(shell go env GOROOT)
VERSION=0.0.2
BIN=genpasswd
DIR_SRC=./cmd/genpasswd

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(GOROOT)/bin/go

build:$(DIR_SRC)/main.go
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

docker_image: clean
	@docker build -f ./Dockerfile -t genpasswd:$(VERSION) .

install: build
	@$(GO) install $(GO_FLAGS) $(DIR_SRC)

test:
	@$(GO) test .

dist: clean
	@goreleaser  --skip-publish --rm-dist --snapshot

release:
	@goreleaser --rm-dist

test-race:
	@$(GO) test -race .

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)
	@rm -rf ./dist/*
