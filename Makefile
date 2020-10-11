.PHONY: all test cover

all: get build test

default: all

get:
	go get ./...

build:
	go build ./...

test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install:
	go install ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

cover: test
	go tool cover

ci-local:
	circleci local execute