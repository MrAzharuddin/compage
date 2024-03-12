package cmd

import (
	"os"

	genInit "github.com/intelops/compage/cmd/genai/init"
	"github.com/intelops/compage/cmd/internal/utils"
	testcmd "github.com/intelops/compage/cmd/testCmd"
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
	testcmd := testcmd.NewTestCmd(accessToken)
	genaiInit := genInit.NewGenAIStart(log, vprConfig)

	rootCmd.AddCommand(testcmd.Execute())
	rootCmd.AddCommand(genaiInit.Execute())

}
