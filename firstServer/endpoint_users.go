package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endPointUserHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)

	case http.MethodGet:
		apiCfg.handleGetUser(w, r)

	case http.MethodDelete:
		apiCfg.handleDeleteUser(w, r)

	case http.MethodPut:
		apiCfg.handleUpdateUser(w, r)

	default:
		respondWithError(w, 404, errors.New("invalid method"))
	}
}

