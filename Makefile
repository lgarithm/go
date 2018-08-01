.PHONY: tools

tools:
	go version
	GOBIN=$(PWD)/bin \
	go install -v ./tools/cmd/...

install:
	go install -v ./tools/cmd/...

clean:
	go clean -cache
