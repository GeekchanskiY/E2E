BIN="./bin"
SRC=$(shell find . -name "*.go")

.PHONY: generate-openapi fmt lint

default: all

all: fmt lint

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

generate-openapi:
	$(info ******************** generating docs ********************)
	swag fmt
	swag init --dir ./cmd,./internal --parseDependency 2
	redocly build-docs docs/swagger.yaml -o docs/doc.html
