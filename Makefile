.PHONY: build build-admin build-account run-admin run-account test

BIN_DIR := bin
ADMIN_BINARY := $(BIN_DIR)/terraform-provider-briaadmin
HOSTNAME=galoymoney


version = 0.1.0
os_arch = $(shell go env GOOS)_$(shell go env GOARCH)
provider_path = registry.terraform.io/galoymoney/briaadmin/$(version)/$(os_arch)/

fmt:
	goimports -l -w .
	go mod tidy
	terraform fmt --recursive

build: build-admin build-regular

build-admin:
	go build -o $(ADMIN_BINARY) cmd/admin-provider/main.go

install: build-admin 
	mkdir -p ~/.terraform.d/plugins/${provider_path}
	mv ${ADMIN_BINARY} ~/.terraform.d/plugins/${provider_path}


testacc:
	TF_ACC=1 go test -v ./...
