all: dependencies test

cover:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out
	@echo
	@echo "'make cover html=true' to see coverage details in a browser"
ifeq ("$(html)","true")
	go tool cover -html=cover.out
endif
	@rm cover.out

dependencies:
	go get
	go get github.com/fzipp/gocyclo
	go get github.com/golang/lint

fmt:
	go fmt -x

reportcard: fmt
	gocyclo -over 10 .
	golint
	go vet

test:
	go test -v -cover ./...

test_failures:
	go test -v ./... 2>&1 | grep -A 1 FAIL
