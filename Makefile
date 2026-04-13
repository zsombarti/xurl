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

# lint requires golangci-lint: https://golangci-lint.run/usage/install/
.PHONY: lint
lint:
	golangci-lint run ./...

# default target: just build locally (format separately to avoid surprise changes)
.PHONY: all
all: build

# ci runs format check, build, and tests with the race detector
.PHONY: ci
ci: lint format build test

.PHONY: release
release:
	goreleaser release --clean
