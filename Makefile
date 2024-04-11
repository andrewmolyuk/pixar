.DEFAULT_GOAL := build
.NOTPARALLEL:
.SILENT:

lint: 
	gocyclo -over 15 .
	golangci-lint run ./...
	staticcheck -checks 'all,-ST*' ./...
	go vet ./...
	markdownlint *.md
.PHONY: lint

test: lint
	go test ./... -cover
.PHONY: test

build: lint test
	go build -ldflags "-s -w" -o ./bin/pixar ./cmd/main.go
.PHONY: build

sim:
	rm -Rf ./testdata/output && go run ./cmd/main.go --input ./testdata/input --output ./testdata/output --policy folder -s
.PHONY: sim

dev:
	rm -Rf ./testdata/output && go run ./cmd/main.go --debug --input ./testdata/input --output ./testdata/output
.PHONY: dev

upgrade:
	go get -u ./...
	go mod tidy
.PHONY: update

