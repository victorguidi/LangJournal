package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Language struct {
	Language       string `json:"language"`
	Level          string `json:"level"`
	TimeOfLearning string `json:"timeOfLearning"`
}

func (a *API) GetLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error) {
	return []Language{}, nil
}

func (a *API) GenerateLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error) {
	return []Language{}, nil
}

func (a *API) UpdateLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error) {
	return []Language{}, nil
}

func (a *API) DetermineLanguageLvl(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Language string `json:"language"`
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Base Questions for determine Language Level, get from the database
	baseQuestions, err := a.db.GetBaseQuestions()
	if err != nil {
		log.Default().Fatalf("Error getting base questions: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	final, err := a.TranslateQuestions(req.Language, baseQuestions)
	if err != nil {
		log.Default().Fatalf("Error translating questions: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(final)
}
