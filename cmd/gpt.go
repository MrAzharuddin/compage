package cmd

import (
	"os"

	genInit "github.com/intelops/compage/cmd/genai/init"
	unitTest "github.com/intelops/compage/cmd/genai/unit-test"
	"github.com/intelops/compage/cmd/internal/utils"
)

func init() {
	// add viper configuration
	vprConfig, err := utils.LoadViper()
	logger := utils.NewLog()
	log := logger.GetLogger()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// create a new testCmd instance
	genaiInit := genInit.NewGenAIStart(log, vprConfig)
	genaiUnitTest := unitTest.NewUnitTestCmd(log, vprConfig)

	// Add SubCommands for gpt
	rootCmd.AddCommand(genaiInit.Execute())
	rootCmd.AddCommand(genaiUnitTest.Execute())

}
