.PHONY: bench cover coverage.txt dependencies fmt reportcard sec test test_failures

all: dependencies test

bench:
	go test -bench .

cover:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out
	@echo
	@echo "'make cover html=true' to see coverage details in a browser"
ifeq ("$(html)","true")
	go tool cover -html=cover.out
endif
	@rm cover.out

coverage.txt:
	go test -v -coverprofile=$@ -covermode=atomic

cover-html:
	$(MAKE) cover html=true

dependencies:
	go get -t -v  ./...
	go get github.com/fzipp/gocyclo
	go get golang.org/x/lint/golint
	go get github.com/securego/gosec/cmd/gosec/...

fmt:
	go fmt -x

reportcard: fmt
	gocyclo -over 10 .
	golint
	go vet

sec:
	gosec ./...

test:
	go test -v -cover ./...

test_failures:
	go test -v ./... 2>&1 | grep -A 1 FAIL
