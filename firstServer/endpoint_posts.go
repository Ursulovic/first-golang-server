package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endPointPostHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)

	case http.MethodGet:
		apiCfg.handlerRetrievePosts(w, r)

	case http.MethodDelete:
		apiCfg.handlerDeletePost(w, r)

	default:
		respondWithError(w, 404, errors.New("invalid method"))
	}
}

