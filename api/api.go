package api

import (
	"encoding/json"
	"net/http"

	"github.com/alextanhongpin/go-openapi/petstore"
)

type FindPetsParams = petstore.FindPetsParams
type Pet = petstore.Pet

type PetStoreImpl struct{}

func (*PetStoreImpl) FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams) {
	var result []Pet
	result = make([]Pet, 1)
	result[0] = Pet{}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*PetStoreImpl) AddPet(w http.ResponseWriter, r *http.Request) {
	var pet Pet

	// Now, we have to return the NewPet
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func (*PetStoreImpl) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	var pet Pet

	// Now, we have to return the NewPet
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)
}

func (*PetStoreImpl) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNoContent)
}
