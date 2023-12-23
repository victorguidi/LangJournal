package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/victorguidi/LangJournaling/src/db"
)

type ExternalApi struct {
	URL   string
	Model string
}

func NewExternalAPI(url string, model string) *ExternalApi {
	return &ExternalApi{
		URL:   url,
		Model: model,
	}
}

// TODO: REFACTOR THIS FUNCTION IS TOO BIG
func (a *API) TranslateQuestions(language string, questions interface{}) (interface{}, error) {

	var questionsToTranslate []string

	for _, q := range questions.([]db.BaseQuestions) {
		questionsToTranslate = append(questionsToTranslate, fmt.Sprintf(`
      question: %s
      answer: %s
      `, q.Question, q.Answer))
	}

	// DEFAULT PROMPT FOR ASK THE MODEL TO TRANSLATE
	prompt := fmt.Sprintf(`YOU ARE A PROFESSIONAL LANGUAGE TEACHER. YOUR TASK IS TO TRANSLATE THE FOLLOWING QUESTIONS AND ANSWERS INTO %s. TRANSLATE THIS: %v`, language, strings.Join(questionsToTranslate, ""))

	type Payload struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
	}

	payloadData := Payload{
		Model:  a.ExternalApi.Model,
		Prompt: prompt,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", a.ExternalApi.URL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	type Response struct {
		Model      string `json:"model"`
		Created_at string `json:"created_at"`
		Response   string `json:"response"`
		Done       bool   `json:"done"`
	}
	var resp []Response

	for _, b := range strings.Split(string(body), "\n") {
		if b != "" {
			var r Response
			err := json.Unmarshal([]byte(b), &r)
			if err != nil {
				log.Default().Fatalf("Error unmarshalling response: %v", err)
				return nil, err
			}
			resp = append(resp, r)
		}
	}

	type FinalAnswer struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	var finalAnswers []FinalAnswer

	// TODO: THIS IS SO BAD, FIX THIS ASAP. RIGHT NOW IT IS ONLY WORKING SOMETIMES, NOT ALWAYS
	for {
		var question string
		currIndex := 0
	first:
		for i, r := range resp {
			if r.Response == ":" {
				for o, r := range resp[i+1:] {
					question += r.Response
					if r.Response == "\n" {
						currIndex = o
						break first
					}
				}
			}
		}
		var answer string
		for _, r := range resp[currIndex+1:] {
			answer += r.Response
			if r.Response == "\n" {
				break
			}
		}
		finalAnswers = append(finalAnswers, FinalAnswer{
			Question: question,
			Answer:   answer,
		})

		if len(finalAnswers) == len(questions.([]db.BaseQuestions)) {
			break
		}
	}

	return finalAnswers, nil
}
