.PHONY: build test

BASE_PATH=$(shell pwd)

build:
		docker run --rm -v $(BASE_PATH):/go/src/github.com/zwwhdls/csibuilder \
    	-v $(BASE_PATH)/bin:/bin/csibuilder \
    	-w /go/src/github.com/zwwhdls/csibuilder \
    	golang:1.18 sh ./hack/multibuild.sh ./cmd /bin/csibuilder

test:
	go test -v -race -cover ./pkg/... -coverprofile=cov.out
