default: build

build:
	go build .

fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

fmt-generated:
	gofmt -w ../../sdk
	goimports -w ../../sdk

run: build
	@echo "==> generating SDKs..."
	./autowrapper
	@echo "==> updating the user agent..."
	../../scripts/update-user-agent.sh

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@latest
	sudo npm install -g autorest
	@echo "==> running 'npm install'..."
	npm install

.PHONY: build fmt fmt-generated run tools
