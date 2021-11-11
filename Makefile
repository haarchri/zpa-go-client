# ====================================================================================
# Setup Project

all: generate

generate:
	swagger generate client -f doc/zpa-api-docs.json --skip-validation -c pkg/client -m pkg/models --default-scheme=https
