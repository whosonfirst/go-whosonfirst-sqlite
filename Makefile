fmt:
	go fmt cmd/wof-sqlite-index-example/*.go
	go fmt database/*.go
	go fmt index/*.go
	go fmt tables/*.go
	go fmt utils/*.go

tools:
	go build -o bin/wof-sqlite-index-example cmd/wof-sqlite-index-example/main.go
