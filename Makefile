.PHONY: gen-proto build build-admin testacc

BIN_OUT_DIR := out
ADMIN_BINARY := $(BIN_OUT_DIR)/terraform-provider-briaadmin
ACCOUNT_BINARY := $(BIN_OUT_DIR)/terraform-provider-briaaccount
HOSTNAME=galoymoney

PROTO_DIR := proto/vendor
PROTO_OUTPUT_DIR := bria/proto

version = 0.1.0
os_arch = $(shell go env GOOS)_$(shell go env GOARCH)
admin_provider_path = registry.terraform.io/galoymoney/briaadmin/$(version)/$(os_arch)/
account_provider_path = registry.terraform.io/galoymoney/briaaccount/$(version)/$(os_arch)/

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


build-admin:
	go build -o $(ADMIN_BINARY) cmd/admin-provider/main.go

build-account:
	go build -o $(ACCOUNT_BINARY) cmd/account-provider/main.go

install: gen-proto build-admin build-account
	mkdir -p ~/.terraform.d/plugins/${admin_provider_path}
	mkdir -p ~/.terraform.d/plugins/${account_provider_path}
	mv ${ADMIN_BINARY} ~/.terraform.d/plugins/${admin_provider_path}
	mv ${ACCOUNT_BINARY} ~/.terraform.d/plugins/${account_provider_path}
	rm -rf example/.terraform example/.terraform.lock.hcl


testacc:
	TF_ACC=1 go test -v ./...
