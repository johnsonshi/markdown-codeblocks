.PHONY: build-cli
build-cli:
	go build -v -o ./bin/md-cb ./cmd/cli
	chmod +x ./bin/md-cb

.PHONY: test
test:
	go test -v ./...

.PHONE: test-update-golden
test-update-golden:
	go test -v ./... -update
