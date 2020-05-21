# Copyright 2019-2020 Axetroy. All rights reserved. Apache License 2.0.

test:
	GO_TESTING=1 go test --cover -covermode=count -coverprofile=coverage.out ./...

test-ci:
	GO_TESTING=1 GITHUB_CI=1 go test --cover -covermode=count -coverprofile=coverage.out ./...

start:
	GO111MODULE=on go run ./main.go start

build:
	bash build.sh