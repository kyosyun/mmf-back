include .env

DOT_ENV := `cat .env | grep -v ^\#|xargs`

.PHONY: devtool
devtool:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: codegen
codegen:
	@$(MAKE) codegen-oapi
	@$(MAKE) codegen-db

.PHONY: codegen-oapi
codegen-oapi:
	oapi-codegen --config api/openapi.models.gen.cfg.yaml api/openapi.yaml > api/openapi.models.gen.go
	oapi-codegen --config api/openapi.server.gen.cfg.yaml api/openapi.yaml > api/openapi.server.gen.go

.PHONY: codegen-db
codegen-db:
	rm -f ./internal/infra/postgres/model/*.gen.go
	env ${DOT_ENV} go run tools/codegen-db-models/main.go --output="./internal/infra/postgres/model"