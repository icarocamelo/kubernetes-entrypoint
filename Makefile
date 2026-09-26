build:
	@echo "Building kubernetes-entrypoint for $(GOOS)/$(GOARCH)"
	mkdir -p bin/$(GOARCH)
	go build -o bin/$(GOARCH)/kubernetes_entrypoint

linux-arm64:
	export GOOS="linux"; \
	export GOARCH="arm64"; \
	$(MAKE) build

linux-amd64:
	export GOOS="linux"; \
	export GOARCH="amd64"; \
	$(MAKE) build

clean:
	rm -rf bin

test:
	go test ./...

all: linux-amd64 linux-arm64
