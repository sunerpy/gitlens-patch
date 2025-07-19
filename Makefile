# 项目变量
IMAGE_NAME = gitlens-patch
DIST_DIR = dist
PLATFORMS = linux/amd64 linux/arm64 windows/amd64 windows/arm64
TAG ?= dev

.PHONY: test test-all build-local build-binaries package-binaries clean lint code-format

# 单元测试
test:
	go test ./...

test-all:
	for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		echo "Testing for $$platform"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go test ./... || exit 1; \
	done

# 本地构建
build-local:
	@echo "Building binary for local environment"
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(IMAGE_NAME) .

# 多平台构建
build-binaries:
	@echo "Building binaries for platforms: $(PLATFORMS)"
	for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		OUTPUT=$(DIST_DIR)/$(IMAGE_NAME)-$$GOOS-$$GOARCH; \
		if [ "$$GOOS" = "windows" ]; then OUTPUT=$$OUTPUT.exe; fi; \
		echo "Building for $$platform -> $$OUTPUT"; \
		GOOS=$$GOOS GOARCH=$$GOARCH CGO_ENABLED=0 go build -o $$OUTPUT . || exit 1; \
	done

# 打包产物
package-binaries:
	@echo "Packaging binaries into tar.gz/zip archives"
	for file in $(DIST_DIR)/$(IMAGE_NAME)-*; do \
		if [[ $$file == *.exe ]]; then \
			zip -j $(DIST_DIR)/$$(basename $$file)-$(TAG).zip $$file; \
		else \
			tar -czvf $(DIST_DIR)/$$(basename $$file)-$(TAG).tar.gz -C $(DIST_DIR) $$(basename $$file); \
		fi; \
	done

# Lint 检查
.PHONY: lint

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...

# 代码格式化
code-format:
	@echo "Running code format..."
	@find . -name "*.go" -type f -exec gofmt -w {} \;

# 清理构建产物
clean:
	rm -rf $(DIST_DIR) 