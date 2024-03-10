PROJECT_NAME := go-year-progress
GOLANGCI_LINT_VER := v1.56.2

.PHONY: build
build:
	 go build -o yp cmd/yp/main.go

.PHONY: release-local
release-local:
	 goreleaser  release --snapshot --skip-publish --clean --rm-dist

.PHONY: go-lint
go-lint:
	docker run \
		--rm \
		--volume "$$PWD":/src/github.com/martinsirbe/$(PROJECT_NAME) \
		--workdir /src/github.com/martinsirbe/$(PROJECT_NAME) \
		golangci/golangci-lint:$(GOLANGCI_LINT_VER) \
		/bin/bash -c "golangci-lint run -v --config=/src/github.com/martinsirbe/$(PROJECT_NAME)/.golangci.yml"
