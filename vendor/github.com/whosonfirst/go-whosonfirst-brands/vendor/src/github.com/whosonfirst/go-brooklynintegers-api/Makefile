prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-brooklynintegers-api; then rm -rf src/github.com/whosonfirst/go-brooklynintegers-api; fi
	mkdir -p src/github.com/whosonfirst/go-brooklynintegers-api
	cp api.go src/github.com/whosonfirst/go-brooklynintegers-api/
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/jeffail/gabs"
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/go-whosonfirst-pool"
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/go-writer-tts"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/int cmd/int.go
	@GOPATH=$(shell pwd) go build -o bin/proxy-server cmd/proxy-server.go

