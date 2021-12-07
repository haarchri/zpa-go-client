# ====================================================================================
# Setup Project

all: generate

generate:
	rm -rf pkg/client/
	rm -rf pkg/models/
	rm -rf README.md
	swagger generate client -f doc/zpa-api-docs.json --skip-validation -c pkg/client -m pkg/models --default-scheme=https
	swagger generate markdown -f doc/zpa-api-docs.json --skip-validation --output readme.md