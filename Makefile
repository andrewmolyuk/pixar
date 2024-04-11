all: build

gocyclo:
	gocyclo -over 15 .

golangci-lint:
	golangci-lint run ./...

staticcheck:
	staticcheck -checks 'all,-ST*' ./...

tidy:
	go mod tidy

fmt: tidy
	go fmt ./...

lint: fmt staticcheck golangci-lint gocyclo

test: lint
	go test ./... -cover

build: lint test
	go build -ldflags "-s -w" -o ./bin/pixar ./cmd/main.go

sim:
	rm -Rf ./testdata/output && go run ./cmd/main.go --input ./testdata/input --output ./testdata/output --policy folder -s

run:
	rm -Rf ./testdata/output && go run ./cmd/main.go --debug --input ./testdata/input --output ./testdata/output

upgrade:
	go get -u ./...
	go mod tidy
.PHONY: update

.NOTPARALLEL:

.PHONY: all gocyclo golangci-lint staticcheck tidy fmt lint test run debug build
