package main

import (
	"encoding/json"
	"net/http"
	"strings"
)


func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}


	user , err := apiCfg.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)

}

func (apiCfg apiConfig) handleDeleteUser(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	email := strings.Trim(url, "/users/")

	err := apiCfg.dbClient.DeleteUser(email)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, 200, struct{}{})

}

func (apiCfg apiConfig) handleUpdateUser(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	email := strings.Trim(url, "/users/")

	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := apiCfg.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)

}

func (apiCfg apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	email := strings.Trim(url, "/users/")

	user , err := apiCfg.dbClient.GetUser(email)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)


}
 

