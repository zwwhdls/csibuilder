LD_FLAGS=-ldflags " \
    -X main.csiBuilderVersion=$(shell git describe --tags --dirty --broken) \
    -X main.goos=$(shell go env GOOS) \
    -X main.goarch=$(shell go env GOARCH) \
    -X main.gitCommit=$(shell git rev-parse HEAD) \
    -X main.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
    "

.PHONY: build
build:
	go build $(LD_FLAGS) -o bin/csibuilder ./cmd/

.PHONY: test
test:
	go test -v -race -cover ./pkg/... -coverprofile=cov.out
