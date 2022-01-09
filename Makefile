SHELL:=/bin/bash
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

define package_bins
	for file in bin/*; do \
		if [[ "$$file" == *".exe"* ]]; then \
			mv $$file cocainate.exe; \
			zip $$file.zip cocainate.exe; \
			short_file_name=cocainate.exe; \
		else \
			mv $$file cocainate; \
			tar -czf $$file.tar.gz cocainate cocainate.8; \
			short_file_name=cocainate; \
		fi; \
		rm $$short_file_name; \
	done
endef

package: build
	$(package_bins)

clean:
	rm -r bin

.PHONY: all build test clean linux mac windows