default: build

build:
	go build .

fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

fmt-generated:
	gofmt -w ../../sdk
	goimports -w ../../sdk

run: build
	./autowrapper
	# update the user agent string
	sed -i '' -e "s/github.com\/Azure\/azure-sdk-for-go\/version/github.com\/tombuildsstuff\/kermit\/version/" -e "s/Azure-SDK-For-Go\//tombuildsstuff\/kermit\//" ../../sdk/**/*.go

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@latest
	sudo npm install -g autorest
	@echo "==> running 'npm install'..."
	npm install

.PHONY: build fmt fmt-generated run tools
