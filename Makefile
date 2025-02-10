export ANSIBLE_TIMEOUT=1

# Have ansible command output status in JSON
# export ANSIBLE_LOAD_CALLBACK_PLUGINS=true
# export ANSIBLE_STDOUT_CALLBACK=json
.PHONY: test build fmt

build:
	go build -o bin/loadek ./cmd/loadek
	go build -o bin/premierleague cmd/premierleague/main.go

fmt:
	find . -name \*.go -type f -print0 | xargs -0 -I{} go fmt {}

clean:
	rm -rf bin/loadek

test:
	go test ./... -coverprofile cover.out -covermode=atomic -coverpkg=./... -v 2>&1
	go tool cover -func cover.out 2>&1

clean-test-cache:
	go clean -testcache

lint:
	golangci-lint run -v

install-dev-deps:
	mkdir -p ~/bin
	wget -q -O ~/go/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64
	chmod u+x ~/go/bin/tailwindcss
	npm install
