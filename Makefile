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

# ci runs lint, format check, build, and tests with the race detector
# note: format runs before build so any fmt changes are caught early
# note: including lint in ci since golangci-lint is available in my environment
# note: skipping lint locally sometimes with `make ci-fast` to save time
.PHONY: ci
ci: format lint build test

# ci-fast skips lint and the race detector for quicker local feedback loops
.PHONY: ci-fast
ci-fast: format build test-fast

# release-dry-run is useful for testing goreleaser config locally without publishing
.PHONY: release-dry-run
release-dry-run:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

# tidy keeps go.mod and go.sum clean; run this before committing dependency changes
.PHONY: tidy
tidy:
	go mod tidy

# check-tidy verifies go.mod and go.sum are up to date (useful in ci to catch forgotten `go mod tidy`)
.PHONY: check-tidy
check-tidy: tidy
	git diff --exit-code go.mod go.sum

# help lists available targets — handy since this Makefile is growing
# personal note: I added descriptions here to make the output more useful
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:' Makefile | grep -v '^#' | sed 's/:.*//' | sort | awk '{printf "  %-20s\n", $$1}'
