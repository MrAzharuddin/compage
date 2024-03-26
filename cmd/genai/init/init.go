package init

import (
	"embed"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//go:embed prompts.json.tmpl
var PromptContentTmpl embed.FS

type GenAIStart struct {
	logger      *logrus.Logger
}

func NewGenAIStart(logger *logrus.Logger) *GenAIStart {
	return &GenAIStart{
		logger:      logger,
	}
}

func (g *GenAIStart) Execute() *cobra.Command {
	genAIInitCmd := &cobra.Command{
		Use:   "genaiInit",
		Short: "`genaiInit` initializes the compage gpt with the OpenAI API key and validates the API KEY if available.",
		Long:  "`compage genaiInit` initializes the compage gpt with the OpenAI API key. Checks for OPENAI_KEY is available in the system environment and validates the API KEY and sends a request for our LLM server which handles its validation with our server.",
		PreRun: func(_ *cobra.Command, _ []string) {
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may need some changes."
			g.logger.Println(yellow + text + reset)
		},
		Run: g.genAIInitCmdRun,
	}

	// Add flags to the command
	// goPrompt, dotnetPrompt, language are the default flags values
	genAIInitCmd.Flags().StringVar(&goPrompt, "goPrompt", goPrompt, "goPrompt")
	genAIInitCmd.Flags().StringVar(&goDocPrompt, "goDocPrompt", goDocPrompt, "goDocPrompt")
	genAIInitCmd.Flags().StringVar(&dotnetPrompt, "dotnetPrompt", dotnetPrompt, "dotnetPrompt")
	genAIInitCmd.Flags().StringVar(&dotnetDocPrompt, "dotnetDocPrompt", dotnetDocPrompt, "dotnetDocPrompt")
	genAIInitCmd.Flags().StringVar(&language, "language", language, "language")

	return genAIInitCmd
}

func (g *GenAIStart) genAIInitCmdRun(cmd *cobra.Command, args []string) {
	// validate the `OPENAI_KEY` from the system environment
	// if not available, throw an error and exit
	// else, continue
	err := g.ValidateOpenAIKey()
	if err != nil {
		g.logger.Error("OpenAI key validation failed")
		return
	}

	// validate the `OPENAI_KEY` with the server
	err = g.CheckOpenAITokens()
	if err != nil {
		g.logger.Error("OpenAI key validation failed")
		return
	}

	err = g.generateConfigFile()
	if err != nil {
		g.logger.Error("Config file generation failed")
		return
	}
	g.logger.Info("Config file generated successfully")
}
