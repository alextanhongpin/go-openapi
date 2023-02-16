gen-types:
	@oapi-codegen -package petstore petstore-expanded.yaml > petstore/petstore.gen.go

gen-server:
	@oapi-codegen -package petstore -generate chi-server petstore-expanded.yaml > api/api.gen.go

install:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
