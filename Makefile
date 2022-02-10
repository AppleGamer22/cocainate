SHELL:=/bin/bash
VERSION:=1.0.2
PACKAGE:="github.com/AppleGamer22/cocainate"
LDFLAGS:=-ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)' -X '$(PACKAGE)/cmd.Hash=$(shell git rev-list -1 HEAD)'"

build: linux mac windows

test:
	@echo "Testing $(VERSION) internal"
	go clean -testcache && go test -v -race -cover ./internal
	@echo "Testing $(VERSION) cmd"
	go clean -testcache && go test -v -race -cover ./cmd

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
	for file in bin/cocainate_$(VERSION)*; do \
		if [[ "$$file" == *".exe"* ]]; then \
			mv $$file cocainate.exe; \
			zip $$file.zip cocainate.exe cocainate.ps1; \
			short_file_name=cocainate.exe; \
		else \
			mv $$file cocainate; \
			tar -czf $$file.tar.gz cocainate cocainate.bash cocainate.fish cocainate.zsh cocainate.1; \
			short_file_name=cocainate; \
		fi; \
		rm $$short_file_name; \
	done
endef

package: build completion
	$(package_bins)
	rm -f cocainate.bash cocainate.fish cocainate.zsh cocainate.ps1

completion:
	go build -v $(LDFLAGS) -o ./bin/cocainate
	./bin/cocainate completion bash > cocainate.bash
	./bin/cocainate completion fish > cocainate.fish
	./bin/cocainate completion zsh > cocainate.zsh
	./bin/cocainate completion powershell > cocainate.ps1
	rm bin/cocainate


clean:
	rm -rf bin cocainate.bash cocainate.fish cocainate.zsh cocainate.ps1

.PHONY: all build test clean package linux mac windows