.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -ldflags="-s -w" -o tmp/toy2 bin/toy2/main.go

build_debug: vet
	go build -gcflags=all="-N -l" -o tmp/toy2 bin/toy2/main.go

tmp:
	mkdir -p tmp

test: tmp
	go test -cover -coverpkg=./... -coverprofile=tmp/cov.out ./...

cov: test
	go tool cover -html=tmp/cov.out
