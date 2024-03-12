package init

import (
	"github.com/intelops/compage/cmd/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GenAIStart struct {
	logger      *logrus.Logger
	viperConfig *utils.ViperConfig
}

func NewGenAIStart(logger *logrus.Logger, viperConfig *utils.ViperConfig) *GenAIStart {
	return &GenAIStart{
		logger:      logger,
		viperConfig: viperConfig,
	}
}

func (g *GenAIStart) Execute() *cobra.Command {
	genAIInitCmd := &cobra.Command{
		Use:   "genaiInit",
		Short: "`genaiInit` initializes the compage gpt with the OpenAI API key and tokenize your authentication.",
		Long:  "`compage genaiInit` initializes the compage gpt with the OpenAI API key and tokenize your authentication. Checks for OPENAI_KEY is available in the system environment and the username in the command flags.",
		PreRun: func(_ *cobra.Command, _ []string) {
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may have some bugs."
			g.logger.Println(yellow + text + reset)
		},
		Run: g.genAIInitCmdRun,
	}

	genAIInitCmd.Flags().StringVar(&USERNAME, "username", "", "User Name")
	return genAIInitCmd
}

func (g *GenAIStart) genAIInitCmdRun(cmd *cobra.Command, args []string) {
	configs, err := g.viperConfig.Unmarshal()
	if err != nil {
		g.logger.Error("Config unmarshal failed")

	}
	if !configs.Validate() {
		g.logger.Error("Config validation failed")

	}

	// validate the `OPENAI_KEY` from the system environment
	// if not available, throw an error and exit
	// else, continue
	err = g.ValidateOpenAIKey()
	if err != nil {
		g.logger.Error("OpenAI key validation failed")

	}

	// validate the `username` from the userName flag
	// if not available, throw an error and exit
	// else, continue
	err = g.ValidateUserName(cmd)
	if err != nil {
		g.logger.Error("User name validation failed", err)

	}

	// validate the tokens
	// if not available, throw an error and exit
	// else, continue
	err = g.CheckOpenAITokens()
	if err != nil {
		g.logger.Error("Token validation failed")
	}

	g.logger.Info("Initialization successful")
}
