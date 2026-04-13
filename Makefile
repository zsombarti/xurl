.PHONY: build
build:
	go build -o xurl

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -f xurl

.PHONY: test
test:
	go test -v -race ./...

# run tests without the race detector for faster local iteration
.PHONY: test-fast
test-fast:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: all
all: format build test

.PHONY: release
release:
	goreleaser release --clean
