
include github.com/tj/make/golang

# Build the static site.
build:
	@static-docs \
	  --title "GitHub Polls" \
		--subtitle "SVG polls you can embed in GitHub issues or readmes." \
		--in docs
.PHONY: build

# Clean build artifacts.
clean:
	@rm -fr build
.PHONY: clean
