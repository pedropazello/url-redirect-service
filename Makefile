.PHONY: mocks clean-mocks

# Generate mocks using Docker
mocks:
	docker run --rm -v $(PWD):/src -w /src vektra/mockery:v2.46.3

# Clean generated mocks
clean-mocks:
	rm -rf mocks/*.go

# Regenerate mocks (clean and generate)
regenerate-mocks: clean-mocks mocks
