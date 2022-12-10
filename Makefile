.PHONY: all test cover
COVER_FILE=./.coverage.txt

all: get build test

default: all

get:
	go get ./...

build:
	go build ./...

test:
	go test ./... -v -coverprofile $(COVER_FILE)
	go tool cover -func $(COVER_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) || true
	rm $(COVER_FILE) || true

install:
	go install ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

cover: test
	go tool cover -func $(COVER_FILE)

ci-local:
	circleci local execute