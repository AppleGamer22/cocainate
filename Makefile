VERSION:=1.0.0

all: test build

build: linux mac windows

test:
	@echo "Testing $(VERSION) internal"
	go clean -testcache && go test -v -cover ./internal/*
	@echo "Testing $(VERSION) cmd"
	go clean -testcache && go test -v -cover ./cmd/*

linux:
	env GOOS=linux GOARCH=amd64 go build -v -o ./bin/cocainate_$(VERSION)_linux_amd64
	env GOOS=linux GOARCH=arm64 go build -v -o ./bin/cocainate_$(VERSION)_linux_arm64
	env GOOS=linux GOARCH=riscv64 go build -v -o ./bin/cocainate_$(VERSION)_linux_riscv64

mac:
	env GOOS=darwin GOARCH=amd64 go build -v -o ./bin/cocainate_$(VERSION)_mac_amd64
	env GOOS=darwin GOARCH=arm64 go build -v -o ./bin/cocainate_$(VERSION)_mac_arm64

windows:
	env GOOS=windows GOARCH=amd64 go build -v -o ./bin/cocainate_$(VERSION)_windows_amd64.exe
	env GOOS=windows GOARCH=arm64 go build -v -o ./bin/cocainate_$(VERSION)_windows_arm64.exe

.PHONY: all build test linux mac windows