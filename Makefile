.PHONY: bench cover coverage.txt dependencies fmt reportcard sec test test_failures

all: dependencies test

bench:
	gotest -bench .

cover:
	gotest -v -coverprofile=cover.out
	go tool cover -func=cover.out
	@echo
	@echo "'make cover html=true' to see coverage details in a browser"
ifeq ("$(html)","true")
	go tool cover -html=cover.out
endif
	@rm cover.out

coverage.txt:
	gotest -v -coverprofile=$@ -covermode=atomic

cover-html:
	$(MAKE) cover html=true

dependencies:
	go get -t -v  ./...
	go get github.com/fzipp/gocyclo
	go get golang.org/x/lint/golint
	go get github.com/securego/gosec/cmd/gosec/...
	go get -u github.com/rakyll/gotest

fmt:
	go fmt -x

reportcard: fmt
	gocyclo -over 10 .
	golint
	go vet

sec:
	gosec ./...

test:
	gotest -v -cover ./...

test_failures:
	gotest -v ./... 2>&1 | grep -A 1 FAIL

test-run:
ifdef test
	gotest -i ./...
	gotest -v -failfast ./... -run $(test)
else
	@echo Syntax is 'make $@ test=<test name>'
endif
