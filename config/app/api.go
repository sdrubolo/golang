package main

import (
	"encoding/json"
	"net/http"

	gocache "github.com/patrickmn/go-cache"
)

type ApiHandler struct {
	store *gocache.Cache
}

func (api *ApiHandler) list(w http.ResponseWriter, _ *http.Request) {
	fileContent, found := api.store.Get("file-content")

	if !found {
		// If there is an error, respond with a 500 status code
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(fileContent)
	if err != nil {
		// If there is an error, respond with a 500 status code
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
