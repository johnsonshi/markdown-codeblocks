.PHONY: build-cli
build-cli:
	go build -v -o ./bin/markdown-codeblocks ./cmd/cli
	chmod +x ./bin/markdown-codeblocks

.PHONY: test
test:
	go test -v ./...

.PHONE: test-update-golden
test-update-golden:
	go test -v ./... -update
