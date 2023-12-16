package api

import (
	db "github.com/victorguidi/LangJournaling/src/db"
	"net/http"
)

type IAPI interface {
	GetPreviousChats(w http.ResponseWriter, r *http.Request) ([]Chat, error)
	GetProgress(w http.ResponseWriter, r *http.Request) ([]Chat, error)
	GenerateChat(w http.ResponseWriter, r *http.Request) ([]Chat, error)
	UpdateChat(w http.ResponseWriter, r *http.Request) ([]Chat, error)
	DeleteChat(w http.ResponseWriter, r *http.Request) ([]Chat, error)
	GetUserInfo(w http.ResponseWriter, r *http.Request) ([]User, error)
	UpdateUserInfo(w http.ResponseWriter, r *http.Request) ([]User, error)
	GetLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error)
	GenerateLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error)
	DetermineLanguageLvl(w http.ResponseWriter, r *http.Request) ([]Language, error)
	UpdateLanguageInfo(w http.ResponseWriter, r *http.Request) ([]Language, error)
}

// This API should implement the http.Handler interface:
type API struct {
	ListenAddr string
	db         *db.DB
	IAPI
}

func New(listenAddr string) *API {
	db, err := db.New()
	if err != nil {
		return nil
	}

	return &API{
		ListenAddr: listenAddr,
		db:         db,
	}
}
