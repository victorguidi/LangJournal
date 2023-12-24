package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
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

	for _, q := range questions.([]db.BaseQuestions) {
		questionsToTranslate = append(questionsToTranslate, fmt.Sprintf(`
      question: %s
      answer: %s
      `, q.Question, q.Answer))
	}

	prompt := fmt.Sprintf(`YOU ARE A PROFESSIONAL LANGUAGE TEACHER. YOUR TASK IS TO TRANSLATE THE FOLLOWING QUESTIONS AND ANSWERS INTO %s. TRANSLATE THIS: %v`, language, strings.Join(questionsToTranslate, ""))
	llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel("mistral")))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Content: prompt},
		schema.HumanChatMessage{Content: "Translate the questions please"},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = completion

	return nil, nil
}
