package docgeneration

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	excludedDirs []string
	language     = "go"
)

// DocGenerationCmd is the struct for the genaiDocGeneration command
type DocGenerationCmd struct {
	logger *logrus.Logger
}

// NewDocGenerationCmd creates a new instance of DocGenerationCmd
func NewDocGenerationCmd(logger *logrus.Logger) *DocGenerationCmd {
	return &DocGenerationCmd{logger: logger}
}

// Execute runs the cobra command for the genaiDocGeneration command
func (dg *DocGenerationCmd) Execute() *cobra.Command {
	// create a new cobra command for the genaiDocGeneration command
	docGenerationCmd := &cobra.Command{
		Use:   "genaiDocGeneration",
		Short: "Generates documentation for the your projects using compage llm",
		Long:  "Generates documentation for the your projects using compage llm and writes it in the current working directory", // Generate AI Documentation
		PreRun: func(_ *cobra.Command, _ []string) {
			// Warn the user that the command is in alpha version and may have some bugs
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may have some bugs."
			dg.logger.Println(yellow + text + reset)
		},
		Run: dg.docGenerationCmdRun, // Set the Run function to docGenerationCmdRun
	}

	docGenerationCmd.Flags().StringSliceVar(&excludedDirs, "excludedDirs", excludedDirs, "excludedDirs")
	docGenerationCmd.Flags().StringVar(&language, "language", language, "language")

	return docGenerationCmd
}

// docGenerationCmdRun is the actual function that runs when the command is executed
func (dg *DocGenerationCmd) docGenerationCmdRun(cmd *cobra.Command, args []string) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		dg.logger.Error("Error getting current working directory:", err)
		return
	}

	// Fetch folders data
	result, err := dg.fetchFoldersData(cwd)
	if err != nil {
		dg.logger.Error("Error fetching folders data:", err)
		return
	}

	folderStructure := strings.Join(result, " --> ")

	response, err := dg.docGenerationServer(folderStructure)
	if err != nil {
		dg.logger.Error("Error generating AI documentation:", err)
		return
	}

	err = dg.storeDocumentationInFile(cwd, response.Data.Docs)
	if err != nil {
		return
	}

	dg.logger.Println("Documentation generated successfully!")

}
