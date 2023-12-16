package api

import (
	"encoding/json"
	"net/http"
)

type Language struct {
	Language       string `json:"language"`
	Level          string `json:"level"`
	TimeOfLearning string `json:"timeOfLearning"`
}

type BaseQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
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
	baseQuestions, err := a.db.GetBaseQuestions(req.Language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if questions, ok := baseQuestions.([]BaseQuestion); ok {
		// Send the base questions and answers to the LLM and translate them into the user's selected language
		// Send the translated questions and answers to the user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
		// w.WriteHeader(http.StatusInternalServerError)
	}
}
