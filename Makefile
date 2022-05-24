GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find app/*/*/internal -name *.proto)
CONFIG_PROTO_FILES=$(shell find config -name *.proto)
API_PROTO_FILES=$(shell find proto -name *.proto)

.PHONY: init
# Init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/envoyproxy/protoc-gen-validate
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/google/wire/cmd/wire
	go mod tidy

.PHONY: errors
# Generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config
# Generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
		   --validate_out=paths=source_relative,lang=go:. \
	       $(CONFIG_PROTO_FILES)

.PHONY: swagger
# Generate api swagger
swagger:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --openapiv2_out . \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   $(API_PROTO_FILES)

.PHONY: api
# Generate api proto
api: swagger errors
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
		   --validate_out=paths=source_relative,lang=go:. \
	       $(API_PROTO_FILES)

.PHONY: release	
# Build release
release:
	@goreleaser release --snapshot --rm-dist

# Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
