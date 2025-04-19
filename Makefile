# Makefile
.PHONY: run build install

.PHONY: dev

dev:
	air

run:
	go run cmd/api/main.go

build:
	go build -o bin/app cmd/api/main.go

install:
	go install ./cmd/api/

clean:
	rm -rf bin/