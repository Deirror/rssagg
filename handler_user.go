package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Deirror/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(wr http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Error pasrsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusOK, databaseUserToUser(user))
}
