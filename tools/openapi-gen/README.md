# OPENAPI GENERATOR

```bash
task gen-openapi

# or
rm -rf internal/generated/openapi
mkdir -p internal/generated/openapi
go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --config tools/openapi-gen/config.yaml api/api-spec.yaml   
go mod tidy
```