.PHONY: build-cli
build-cli:
	go build -v -o ./bin/md-cb ./cmd/cli
	chmod +x ./bin/md-cb
