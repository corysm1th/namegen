.PHONY: build lint test

build:
	GO111MODULE=on \
	go build cmd/namegen.go

lint:
	golangci-lint run

test:
	GO111MODULE=on \
	ginkgo ./...
