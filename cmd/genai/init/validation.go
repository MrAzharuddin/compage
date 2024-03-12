package init

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var (
	OPENAI_KEY string
	USERNAME   string
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

func (ga *GenAIStart) ValidateUserName(cmd *cobra.Command) error {
	var err error
	USERNAME, err = cmd.Flags().GetString("username")
	if err != nil {
		return err
	}
	if USERNAME == "" {
		return errors.New("`username` is not set in the command flags")
	}
	return nil
}
