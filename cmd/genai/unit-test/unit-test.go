package unitTest

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/intelops/compage/cmd/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	config *models.PromptsConfig
	err error
)

// UnitTestCmd is the struct for the genaiUnitTest command
type UnitTestCmd struct {
	logger *logrus.Logger
}

// NewUnitTestCmd returns a new instance of UnitTestCmd
func NewUnitTestCmd(logger *logrus.Logger) *UnitTestCmd {
	// Create a new UnitTestCmd struct with the given logger
	return &UnitTestCmd{
		logger: logger,
	}
}

// Execute returns the cobra command for the genaiUnitTest command
func (u *UnitTestCmd) Execute() *cobra.Command {
	// Create a new cobra command for the genaiUnitTest command
	unitTestCmd := &cobra.Command{
		Use:   "genaiUnitTest",
		Short: "Generate Unit Test using compage llm",                                                                                                                                                                             // Generate AI Unit Test
		Long:  `Generate Unit Test using compage llm backend server which handles its validation and authentication with our server. It will generate unit tests for your code and stores them in the current working directory.`, // Generate AI Unit Test
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
	// set config
	config, err = models.ReadPromptsConfigJSONFile(utils.PROMPTS_CONFIG_FILE)
	if err != nil {
		u.logger.Error("Error reading prompts config:", err)
		return
	}
	
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		u.logger.Error("Error getting current working directory:", err)
		return
	}

	// Walk through the directory tree and process Go files
	u.logger.Info("Processing files in directory tree:", cwd)
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			u.logger.Error("Error accessing path:", err)
			return nil
		}

		// Only process regular files
		if info.Mode().IsRegular() {
			data, err := u.fileProcess(path)
			if err != nil {
				u.logger.Error("Error processing file:", err)
				return nil
			}

			if data == nil {
				return nil
			}

			for _, v := range data {
				if len(v) > 0 {
					codeBlock := strings.Join(v, "\n")
					// Fetch the unit test from OpenAI
					res, err := u.FetchUnitTestFromOpenAI(codeBlock)
					if err != nil {
						u.logger.Error("Error fetching unit test from OpenAI:", err)
						return nil
					}
					// process the file information properly
					relativePath, err := filepath.Rel(cwd ,path)
					if err != nil {
						u.logger.Error("Error getting relative path:", err)
						return nil
					}
					u.logger.Info(relativePath)
					dirPath := filepath.Dir(relativePath)
					u.logger.Info("Directory path:",dirPath)
					err = u.processCodeBlock(relativePath, res.Data.Code)
					if err != nil {
						u.logger.Error("Error processing code block:", err)
						return nil
					}
					u.logger.Info("Generated unit test for " + relativePath)
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
