.PHONY: build-debug
build-debug:
	go mod download
	type dlv || go install github.com/go-delve/delve/cmd/dlv@v1.8.0
	mkdir -p tmp
	go build -gcflags "all=-N -l" -buildvcs=false -o ./tmp/graphql-app ./graph/server