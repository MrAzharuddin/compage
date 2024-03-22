package unitTest

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UnitTestCmd is the struct for the genaiUnitTest command
type UnitTestCmd struct {
	logger      *logrus.Logger
}

// NewUnitTestCmd returns a new instance of UnitTestCmd
func NewUnitTestCmd(logger *logrus.Logger) *UnitTestCmd {
	// Create a new UnitTestCmd struct with the given logger and viperConfig
	return &UnitTestCmd{
		logger:      logger,
	}
}

// Execute returns the cobra command for the genaiUnitTest command
func (u *UnitTestCmd) Execute() *cobra.Command {
	// Create a new cobra command for the genaiUnitTest command
	unitTestCmd := &cobra.Command{
		Use:   "genaiUnitTest",
		Short: "Generate AI Unit Test", // Generate AI Unit Test
		Long:  `Generate AI Unit Test`, // Generate AI Unit Test
		PreRun: func(_ *cobra.Command, _ []string) {
			// Warn the user that the command is in alpha version and may have some bugs
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may have some bugs."
			u.logger.Println(yellow + text + reset)
		},
		Run: u.unitTestCmdRun, // Set the Run function to unitTestCmdRun
	}
	return unitTestCmd
}

// unitTestCmdRun is the actual function that runs when the command is executed
// It walks through the current working directory and its subdirectories,
// and generates unit tests for all the functions in Go files.
func (u *UnitTestCmd) unitTestCmdRun(cmd *cobra.Command, args []string) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		u.logger.Error("Error getting current working directory:", err)
		return
	}

	// Walk through the directory tree and process Go files
	u.logger.Info("Processing Go files in directory tree:", cwd)
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			u.logger.Error("Error accessing path:", err)
			return nil
		}

		// Only process regular files with a .go extension
		if info.Mode().IsRegular() && filepath.Ext(path) == ".go" {
			data, err := u.fileProcess(path)
			if err != nil {
				u.logger.Error("Error processing Go file:", err)
				return nil
			}

			if data == nil {
				return nil
			}

			for k, v := range data {
				if len(v) > 0 {
					res, err := u.FetchUnitTestFromOpenAI(strings.Join(v, "\n"))
					if err != nil {
						u.logger.Error("Error fetching unit test from OpenAI:", err)
						return nil
					}
					err = u.processCodeBlock(k, res.Data.Code)
					if err != nil {
						u.logger.Error("Error processing code block:", err)
						return nil
					}
					u.logger.Info("Generated unit test for " + k)
				}
			}
		}
		return nil
	})

	if err != nil {
		u.logger.Error("Error walking through directory:", err)
		return
	}
}
