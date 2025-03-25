package resources

import (
	"errors"
	"os"
)

type Env struct {
	OpenAIAPIKey string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Load() error {
	key, ok := os.LookupEnv("OPENAI_API_KEY"); if !ok {
		return errors.New("OPENAI_API_KEY is not set")
	}
	e.OpenAIAPIKey = key
	return nil
}

func (e *Env) GetOpenAIAPIKey() string {
	return e.OpenAIAPIKey
}
