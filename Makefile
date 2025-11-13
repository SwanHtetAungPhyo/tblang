.PHONY: all build install clean test release

VERSION ?= 0.1.0
LDFLAGS := -s -w -X main.version=$(VERSION)

all: build

build: build-core build-plugins

build-core:
	@echo "Building TBLang core..."
	cd core && go build -ldflags="$(LDFLAGS)" -o tblang ./cmd/tblang

build-plugins: build-aws-plugin

build-aws-plugin:
	@echo "Building AWS provider plugin..."
	cd plugin/aws && go build -o tblang-provider-aws main.go

install: build
	@echo "Installing TBLang..."
	sudo cp core/tblang /usr/local/bin/
	sudo mkdir -p /usr/local/lib/tblang/plugins
	sudo cp plugin/aws/tblang-provider-aws /usr/local/lib/tblang/plugins/
	@echo "✓ TBLang installed successfully!"
	@echo "  Core: /usr/local/bin/tblang"
	@echo "  AWS Plugin: /usr/local/lib/tblang/plugins/tblang-provider-aws"

uninstall:
	@echo "Uninstalling TBLang..."
	sudo rm -f /usr/local/bin/tblang
	sudo rm -rf /usr/local/lib/tblang
	@echo "✓ TBLang uninstalled"

clean:
	@echo "Cleaning build artifacts..."
	rm -f core/tblang
	rm -f plugin/aws/tblang-provider-aws
	rm -rf dist/
	@echo "✓ Clean complete"

test:
	@echo "Running tests..."
	cd core && go test ./...
	cd plugin/aws && go test ./...

release: clean
	@echo "Building release binaries..."
	mkdir -p dist
	
	@echo "Building for macOS (amd64)..."
	cd core && GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ../dist/tblang-darwin-amd64 ./cmd/tblang
	cd plugin/aws && GOOS=darwin GOARCH=amd64 go build -o ../../dist/tblang-provider-aws-darwin-amd64 main.go
	
	@echo "Building for macOS (arm64)..."
	cd core && GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ../dist/tblang-darwin-arm64 ./cmd/tblang
	cd plugin/aws && GOOS=darwin GOARCH=arm64 go build -o ../../dist/tblang-provider-aws-darwin-arm64 main.go
	
	@echo "Building for Linux (amd64)..."
	cd core && GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ../dist/tblang-linux-amd64 ./cmd/tblang
	cd plugin/aws && GOOS=linux GOARCH=amd64 go build -o ../../dist/tblang-provider-aws-linux-amd64 main.go
	
	@echo "Building for Linux (arm64)..."
	cd core && GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ../dist/tblang-linux-arm64 ./cmd/tblang
	cd plugin/aws && GOOS=linux GOARCH=arm64 go build -o ../../dist/tblang-provider-aws-linux-arm64 main.go
	
	@echo "Creating archives..."
	cd dist && tar -czf tblang-$(VERSION)-darwin-amd64.tar.gz tblang-darwin-amd64 tblang-provider-aws-darwin-amd64
	cd dist && tar -czf tblang-$(VERSION)-darwin-arm64.tar.gz tblang-darwin-arm64 tblang-provider-aws-darwin-arm64
	cd dist && tar -czf tblang-$(VERSION)-linux-amd64.tar.gz tblang-linux-amd64 tblang-provider-aws-linux-amd64
	cd dist && tar -czf tblang-$(VERSION)-linux-arm64.tar.gz tblang-linux-arm64 tblang-provider-aws-linux-arm64
	
	@echo "Generating checksums..."
	cd dist && shasum -a 256 *.tar.gz > checksums.txt
	
	@echo "✓ Release build complete!"
	@echo "  Artifacts in: dist/"

dev-setup:
	@echo "Setting up development environment..."
	cd core && go mod download
	cd plugin/aws && go mod download
	@echo "✓ Development environment ready"

fmt:
	@echo "Formatting code..."
	cd core && go fmt ./...
	cd plugin/aws && go fmt ./...
	@echo "✓ Code formatted"

lint:
	@echo "Running linters..."
	cd core && go vet ./...
	cd plugin/aws && go vet ./...
	@echo "✓ Linting complete"

help:
	@echo "TBLang Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build          Build core and plugins"
	@echo "  make install        Install TBLang system-wide"
	@echo "  make uninstall      Uninstall TBLang"
	@echo "  make clean          Clean build artifacts"
	@echo "  make test           Run tests"
	@echo "  make release        Build release binaries for all platforms"
	@echo "  make dev-setup      Setup development environment"
	@echo "  make fmt            Format code"
	@echo "  make lint           Run linters"
	@echo "  make help           Show this help message"
