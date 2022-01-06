VERSION:=1.0.0
PACKAGE:="github.com/AppleGamer22/cocainate"
LDFLAGS:=-ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)' -X '$(PACKAGE)/cmd.Hash=$(shell git rev-list -1 HEAD)'"

all: test build

build: linux mac windows

test:
	@echo "Testing $(VERSION) internal"
	go clean -testcache && go test -v -cover ./internal/*
	@echo "Testing $(VERSION) cmd"
	go clean -testcache && go test -v -cover ./cmd/*

linux:
	GOOS=linux GOARCH=amd64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_linux_amd64
	GOOS=linux GOARCH=arm64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_linux_arm64
	GOOS=linux GOARCH=riscv64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_linux_riscv64

mac:
	GOOS=darwin GOARCH=amd64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_mac_amd64
	GOOS=darwin GOARCH=arm64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_mac_arm64

windows:
	GOOS=windows GOARCH=amd64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_windows_amd64.exe
	GOOS=windows GOARCH=arm64 go build -v $(LDFLAGS) -o ./bin/cocainate_$(VERSION)_windows_arm64.exe

clean:
	rm -r bin

.PHONY: all build test clean linux mac windows