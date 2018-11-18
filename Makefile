.PHONY: tools

CMDS = \
	./tools/cmd/... \
	./rtd/cmd/... \
	./travisci/cmd/...

tools:
	go version
	GOBIN=$(PWD)/bin \
	go install -v $(CMDS)

install:
	go install -v $(CMDS)

clean:
	go clean -cache
