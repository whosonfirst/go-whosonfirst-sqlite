CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test ! -d src; then mkdir src; fi
	if test ! -d src/github.com/whosonfirst/go-writer-tts/; then mkdir -p src/github.com/whosonfirst/go-writer-tts/; fi
	if test ! -d src/github.com/whosonfirst/go-writer-tts/speakers; then mkdir -p src/github.com/whosonfirst/go-writer-tts/speakers; fi
	cp tts.go src/github.com/whosonfirst/go-writer-tts/
	cp speakers/*.go src/github.com/whosonfirst/go-writer-tts/speakers/
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/everdev/mack"

vendor-deps: deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/speak cmd/speak.go
	@GOPATH=$(GOPATH) go build -o bin/read cmd/read.go

fmt:
	go fmt *.go
	go fmt speakers/*.go
