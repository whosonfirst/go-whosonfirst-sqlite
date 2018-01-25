prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-pool; then rm -rf src/github.com/whosonfirst/go-whosonfirst-pool; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-pool
	cp *.go src/github.com/whosonfirst/go-whosonfirst-pool/

deps:   self

fmt:
	go fmt *.go

bin: 	self

