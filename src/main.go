package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorguidi/LangJournaling/src/api"
)

func main() {

	api := api.New(":5000")

	http.HandleFunc("/test", api.DetermineLanguageLvl)

	log.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
