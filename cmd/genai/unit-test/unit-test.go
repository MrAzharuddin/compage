package unitTest

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/intelops/compage/cmd/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UnitTestCmd struct {
	logger      *logrus.Logger
	viperConfig *utils.ViperConfig
}

func NewUnitTestCmd(logger *logrus.Logger, viperConfig *utils.ViperConfig) *UnitTestCmd {
	return &UnitTestCmd{
		logger:      logger,
		viperConfig: viperConfig,
	}
}

func (u *UnitTestCmd) Execute() *cobra.Command {
	unitTestCmd := &cobra.Command{
		Use:   "genaiUnitTest",
		Short: "Generate AI Unit Test",
		Long:  `Generate AI Unit Test`,
		PreRun: func(_ *cobra.Command, _ []string) {
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may have some bugs."
			u.logger.Println(yellow + text + reset)
		},
		Run: u.unitTestCmdRun,
	}
	return unitTestCmd
}

func (u *UnitTestCmd) unitTestCmdRun(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		u.logger.Error("Error getting current working directory:", err)
		return
	}

	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			u.logger.Error("Error accessing path:", err)
			return nil
		}
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
