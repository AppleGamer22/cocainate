SHELL:=/bin/bash
VERSION:=$(shell git describe --tags --abbrev=0)
HASH:=$(shell git rev-list -1 HEAD)
PACKAGE:=github.com/AppleGamer22/cocainate
LDFLAGS:=-ldflags="-X '$(PACKAGE)/cmd.Version=$(subst v,,$(VERSION))' -X '$(PACKAGE)/cmd.Hash=$(HASH)'"

test:
	go clean -testcache
	go test -v -race -cover ./session ./ps ./cmd

debug:
	go build -race $(LDFLAGS) .

completion:
	go run . completion bash > cocainate.bash
	go run . completion fish > cocainate.fish
	go run . completion zsh > cocainate.zsh
	go run . completion powershell > cocainate.ps1

manual:
	sed -i "s/vVERSION/$(VERSION)/" cocainate.1
	sed -i "s/DATE/$(shell date -Idate)/" cocainate.1

clean:
	rm -rf cocainate bin dist cocainate.bash cocainate.fish cocainate.zsh cocainate.ps1
	go clean -testcache -cache

.PHONY: debug test clean completion manual