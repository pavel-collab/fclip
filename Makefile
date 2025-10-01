APP_NAME = fclip
BIN_DIR = bin

.PHONY: all build clean test debug release install

all: build

## Build debug binary (with symbols)
debug:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

## Build optimized release
release:
	mkdir -p $(BIN_DIR)
	go build -ldflags "-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

## Run tests
test:
	go test ./... -v

install:
	@echo "Installing"
	go install ./cmd/fclip/...

## Clean artifacts
clean:
	rm -rf $(BIN_DIR)

## Install binary to GOPATH/bin (or GOBIN if set)
install: release
	@if [ -n "$$GOBIN" ]; then \
	  dest="$$GOBIN"; \
	else \
	  dest="$$GOPATH/bin"; \
	fi; \
	mkdir -p "$$dest" && install -m 0755 $(BIN_DIR)/$(APP_NAME) "$$dest/$(APP_NAME)" && \
	echo "Установлено: $$dest/$(APP_NAME)"