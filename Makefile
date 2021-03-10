.PHONY: build test lint format

.DEFAULT:
build: test
	@goreleaser release --snapshot --rm-dist --skip-publish

test:
	@go test -mod=vendor --cover -covermode=count -coverprofile=coverage.out ./...

lint:
	@golangci-lint run ./... -v

.ONESHELL:
format:
	@gofmt -l -e **/*.go
	@gofmt -l -e *.go
	@go fmt -mod=vendor ./...