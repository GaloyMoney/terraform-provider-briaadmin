.PHONY: gen-proto build build-admin testacc

BIN_DIR := bin
ADMIN_BINARY := $(BIN_DIR)/terraform-provider-briaadmin
HOSTNAME=galoymoney

PROTO_DIR := proto/vendor
PROTO_OUTPUT_DIR := bria/proto

version = 0.1.0
os_arch = $(shell go env GOOS)_$(shell go env GOARCH)
provider_path = registry.terraform.io/galoymoney/briaadmin/$(version)/$(os_arch)/

fmt:
	goimports -l -w .
	go mod tidy
	terraform fmt --recursive

gen-proto:
	mkdir -p $(PROTO_OUTPUT_DIR)
	protoc -I $(PROTO_DIR) \
		--go_out=$(PROTO_OUTPUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUTPUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/admin/api.proto $(PROTO_DIR)/api/bria.proto


build: gen-proto build-admin

build-admin:
	go build -o $(ADMIN_BINARY) cmd/admin-provider/main.go

install: build-admin 
	mkdir -p ~/.terraform.d/plugins/${provider_path}
	mv ${ADMIN_BINARY} ~/.terraform.d/plugins/${provider_path}


testacc:
	TF_ACC=1 go test -v ./...
