SHELL:=/bin/bash
VERSION:=$(shell echo ${$(git describe --tags --abbrev=0)/v})
PACKAGE:=github.com/AppleGamer22/cocainate
LDFLAGS:=-ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)' -X '$(PACKAGE)/cmd.Hash=$(shell git rev-list -1 HEAD)'"

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


clean:
	rm -rf cocainate bin dist cocainate.bash cocainate.fish cocainate.zsh cocainate.ps1
	go clean -testcache -cache

.PHONY: test clean completion debug