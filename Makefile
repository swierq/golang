export ANSIBLE_TIMEOUT=1

# Have ansible command output status in JSON
# export ANSIBLE_LOAD_CALLBACK_PLUGINS=true
# export ANSIBLE_STDOUT_CALLBACK=json
.PHONY: test build fmt

build:
	go build -o bin/loadek cmd/loadek/main.go
	go build -o bin/premierleague cmd/premierleague/main.go

fmt:
	go fmt

clean:
	rm -rf bin/loadek


test:
	go test ./... -coverprofile cover.out -v 2>&1
	go tool cover -func cover.out 2>&1

clean-test-cache:
	go clean -testcache
