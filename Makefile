# Makefile for build Go Mocker

test:
	go test -v --cover -covermode=count -coverprofile=coverage.out ./...

start:
	GO111MODULE=on go run ./cmd/gmock/main.go start

build:
	bash build.sh