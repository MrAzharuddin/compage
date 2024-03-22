package cmd

import (
	genInit "github.com/intelops/compage/cmd/genai/init"
	unitTest "github.com/intelops/compage/cmd/genai/unit-test"
	"github.com/intelops/compage/cmd/internal/utils"
)

func init() {
	// add logger
	logger := utils.NewLog().GetLogger()

	// create a new testCmd instance
	genaiInit := genInit.NewGenAIStart(logger)
	genaiUnitTest := unitTest.NewUnitTestCmd(logger)

	// Add SubCommands for gpt
	rootCmd.AddCommand(genaiInit.Execute())
	rootCmd.AddCommand(genaiUnitTest.Execute())
	
}
