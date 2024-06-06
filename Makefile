# Makefile for Gophers
.PHONY: all protoc purge run run-air help

# Workspace root directory
WORKSPACE_ROOT := $(shell git rev-parse --show-toplevel)

# Directories
PROTO_DIR := $(WORKSPACE_ROOT)/api/proto
GO_OUT_DIR := $(WORKSPACE_ROOT)/grpc/proto-gen-go
OPENAPI_OUT_DIR := $(WORKSPACE_ROOT)/api/http

# Proto files. Exclude third_party directory
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto' -not -path "$(PROTO_DIR)/third_party/*")

# Proto compile command
PROTOC_CMD := protoc \
	-I $(PROTO_DIR) \
	-I $(PROTO_DIR)/third_party \
	--proto_path=$(PROTO_DIR) \
	--go_out=$(GO_OUT_DIR) \
	--go-grpc_out=$(GO_OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(GO_OUT_DIR) \
	--grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=$(OPENAPI_OUT_DIR) \
	--openapiv2_opt=logtostderr=true \
	--openapiv2_opt=allow_repeated_fields_in_body=true \
	--openapiv2_opt=grpc_api_configuration=$(PROTO_DIR)/gw_mapping.yaml

# Default target
all: protoc build run

# Help target (optional)
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  protoc   Compile proto files"
	@echo "  purge    Delete all generated and temporary files"
	@echo "  run      Run server"
	@echo "  run-air  Run server with air"
	@echo "  help     Display this help message"
	@echo ""

# Build target
build:
	@go build -o $(WORKSPACE_ROOT)/bin/server $(WORKSPACE_ROOT)/cmd/server
	@echo "✅ Server binary built successfully"

# Deps target (optional)
deps:
	@go mod tidy
	@echo "✅ Dependencies updated"

# Protoc target
protoc:
	@$(PROTOC_CMD) $(PROTO_FILES)
	@echo "✅ Proto files compiled successfully"

# Purge target (optional)
purge:
	rm -rf $(GO_OUT_DIR) $(OPENAPI_OUT_DIR)
	rm -f $(WORKSPACE_ROOT)/bin
	@echo "✅ All files deleted successfully"

# Run target (optional)
# This starts the server with the web app enabled
run:
	@go run $(WORKSPACE_ROOT)/cmd/server

# Run target with comstrek/air (optional)
run-air:
	@command -v air >/dev/null 2>&1 || go install github.com/air-verse/air@latest
	@air -c .air.toml

# Setup target (optional only for development purposes)
setup:
	@go mod download
	@go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo "✅ Setup completed successfully"