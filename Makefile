APP_NAME = fclip
BIN_DIR = bin

.PHONY: all build clean test debug release

all: build

## Сборка бинарника в режиме отладки (с символами)
debug:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

## Сборка оптимизированного релиза
release:
	go build -ldflags "-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

## Запуск тестов
test:
	go test ./... -v

install:
	@echo "Installing"
	go install ./cmd/fclip/...

## Очистка
clean:
	rm -rf $(BIN_DIR)