.PHONY: lint fix-lint

lint:
	golangci-lint run

fix-lint:
	golangci-lint run --fix
