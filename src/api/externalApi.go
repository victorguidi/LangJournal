package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	// "github.com/tmc/langchaingo/schema"
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

func (a *API) TranslateQuestions(language string, questions interface{}) (interface{}, error) {

	var questionsToTranslate []string

	// TODO: refactor this so questions do not repeat
	for _, q := range questions.([]db.BaseQuestions) {
		questionsToTranslate = append(questionsToTranslate, fmt.Sprintf(`
      question: %s
      answer: %s
      `, q.Question, q.Answer))
	}

	prompt := fmt.Sprintf(`
	   YOU ARE A PROFESSIONAL TRANSLATOR.
	   YOUR TASK IS TO TRANSLATE THE FOLLOWING QUESTIONS AND ANSWERS INTO %s LANGUAGE.

    BEGIN TRANSLATION==========================================
	   %v
    END TRANSLATION============================================

     `,
		language, strings.Join(questionsToTranslate, ""))

	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, prompt,
		llms.WithTemperature(0.8),
		// llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		// 	// fmt.Print(string(chunk))
		// 	return nil
		// }),
	)
	if err != nil {
		log.Fatal(err)
	}

	final := strings.Split(completion, "\n")
	for _, f := range final {
		fmt.Println(f)
	}

	return nil, nil
}
