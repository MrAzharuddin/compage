package init

import (
	"errors"
	"os"
)

var (
	OPENAI_KEY string
)

func (ga *GenAIStart) ValidateOpenAIKey() error {
	// Check if OPENAI_KEY is set in the environment
	var ok bool
	OPENAI_KEY, ok = os.LookupEnv("OPENAI_KEY")
	if !ok {
		return errors.New("`OPENAI_KEY` is not set in the environment")
	}
	return nil
}
