.PHONY: build build-admin build-account run-admin run-account test

BIN_DIR := bin
ADMIN_BINARY := $(BIN_DIR)/admin-provider
REGULAR_BINARY := $(BIN_DIR)/regular-provider

build: build-admin build-regular

build-admin:
	go build -o $(ADMIN_BINARY) cmd/admin-provider/main.go

build-regular:
	go build -o $(REGULAR_BINARY) cmd/regular-provider/main.go

run-admin: build-admin
	$(ADMIN_BINARY)

run-regular: build-regular
	$(REGULAR_BINARY)

test:
	go test ./... -v
