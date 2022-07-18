package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)


func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	newPost, err := apiCfg.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, newPost)
}

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.Trim(path, "/posts/")

	if id == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("missing post id"))
	}

	err := apiCfg.dbClient.DeletePost(id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})

}

func (apiCfg apiConfig) handlerRetrievePosts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	email := strings.Trim(path, "/posts/")

	if email == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("missing email"))
		return
	}

	posts, err := apiCfg.dbClient.GetPosts(email)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)


}