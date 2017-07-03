
include github.com/tj/make/golang

# Release binaries.
release:
	@goreleaser
.PHONY: release
