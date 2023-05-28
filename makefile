app = quanta

.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build ./...

.PHONY: run
run:
	go run ./cmd/$(app)/main.go

.PHONY: openapihttp
openapihttp:
	oapi-codegen -generate types -o internal/infra/port/openapi/todotypes.go -package openapi api/openapi/quanta.yml
	oapi-codegen -generate chi-server -o internal/infra/port/openapi/todoapi.go -package openapi api/openapi/quanta.yml
	oapi-codegen -generate types -o internal/infra/client/port/openapi/todotypes.go -package openapi api/openapi/quanta.yml
	oapi-codegen -generate client -o internal/infra/client/port/openapi/todoapi.go -package openapi api/openapi/quanta.yml
