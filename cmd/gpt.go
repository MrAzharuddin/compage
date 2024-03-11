package cmd

import (
	"os"

	testcmd "github.com/intelops/compage/cmd/testCmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	// add viper configuration
	v, err := AddGPTConfigForViper()
	if err != nil {
		log.Error(err)
	}
	err = v.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	accessToken := v.GetString("OPENAI_ACCESS_TOKEN")
	if accessToken == "" {
		log.Error("OPENAI_ACCESS_TOKEN is not set")
		os.Exit(1)
	}

	// create a new testCmd instance
	testcmd := testcmd.NewTestCmd(accessToken)

	rootCmd.AddCommand(testcmd.Execute())
}
