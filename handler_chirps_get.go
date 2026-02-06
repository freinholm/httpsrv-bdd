package main

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpRetrieveByID(w http.ResponseWriter, r *http.Request) {
	//ID Extraction
	chirpID, err := uuid.Parse(strings.TrimPrefix(r.URL.Path, "/api/chirps/"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse chirp ID", err)
	}
	
	dbChirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirp", err)
		return
	}
/*
	if !dbChirp {
		respondWithError(w, http.StatusNotFound, "Chirp not found", nil)
	}
*/
	chirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		}

	respondWithJSON(w, http.StatusOK, chirp)
}