package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alextanhongpin/go-openapi/api"
	"github.com/alextanhongpin/go-openapi/petstore"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	var myAPI api.PetStoreImpl

	swagger, err := petstore.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(&myAPI, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
