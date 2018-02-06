CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-brands
	cp brands.go src/github.com/whosonfirst/go-whosonfirst-brands
	cp -r whosonfirst src/github.com/whosonfirst/go-whosonfirst-brands/
	cp -r vendor/* src/

deps:   
	@GOPATH=$(GOPATH) go get -u "github.com/tidwall/gjson"
	# @GOPATH=$(GOPATH) go get -u "github.com/tidwall/pretty"
	# @GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-brooklynintegers-api"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-crawl"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-flags"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-json"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt brands.go
	go fmt cmd/*.go
	go fmt whosonfirst/*.go

bin:	self
	# @GOPATH=$(GOPATH) go build -o bin/wof-brands-create cmd/wof-brands-create.go
	@GOPATH=$(GOPATH) go build -o bin/wof-brands-find cmd/wof-brands-find.go
	@GOPATH=$(GOPATH) go build -o bin/wof-brands-crawl cmd/wof-brands-crawl.go
