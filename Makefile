# Makefile
.PHONY: run build install

.PHONY: dev

.PHONY: setup-config run clean

# Автоматический выбор конфига windows / linux
setup-config:

ifeq ($(OS),Windows_NT)
	@if not exist .air-windows.toml ( \
		echo "Error: .air-windows.toml not found" && exit 1 \
	) else ( \
		copy /Y .air-windows.toml .air.toml && \
		echo Windows config applied \
	)
else
	@[ ! -f .air-mac.toml ] && echo "Error: .air-mac.toml not found" && exit 1 || true
	@ln -sf .air-mac.toml .air.toml
	@echo "macOS config applied"
endif

# Запуск dev mode
dev: setup-config
	@echo "Starting air..."
	@air

build:
	go build -o bin/app cmd/api/main.go

install:
	go install ./cmd/api/

clean:
	rm -rf bin/