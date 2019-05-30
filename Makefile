.PHONY: build clean test test-race

VERSION=0.0.2
BIN=genpasswd
DIR_SRC=./cmd/genpasswd

GO_ENV=CGO_ENABLED=1
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=env $(GO_ENV) go

build:$(DIR_SRC)/main.go
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

docker_image: clean
	@docker build -f ./Dockerfile -t genpasswd:$(VERSION) .

install: build
	@$(GO) install $(GO_FLAGS) $(DIR_SRC)

test:
	@$(GO) test .

test-race:
	@$(GO) test -race .

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)
