gen-types:
	@oapi-codegen -package petstore -generate types,client,spec,chi-server petstore-expanded.yaml > petstore/petstore.gen.go

install:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

curl:
	@curl -H 'Content-Type: application/json' 0.0.0.0:8080/pets
