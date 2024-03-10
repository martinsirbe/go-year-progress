.PHONY: build
build:
	 go build -o yp cmd/yp/main.go

.PHONY: release-local
release-local:
	 goreleaser  release --snapshot --skip-publish --clean --rm-dist
