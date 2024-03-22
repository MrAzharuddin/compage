package init

import (
	"fmt"
	"os"

	"github.com/intelops/compage/internal/languages/executor"
)

func (g *GenAIStart) generateConfigFile() error {
	promptFilePath := "prompts.yaml"
	// prompt file does not exist, create it
	_, err := os.Create(promptFilePath)
	if err != nil {
		g.logger.Errorf("error while creating the prompt file %s", err)
		err = os.Remove(promptFilePath)
		if err != nil {
			g.logger.Errorf("error while removing the prompt file %s", err)
			return err
		}
	}

	if language != "go" && language != "dotnet" {
		message := fmt.Sprintf("language %s is not supported", language)
		g.logger.Errorf(message)
		return fmt.Errorf(message)
	}

	if language == "go" {
		promptContentData, err := PromptContentTmpl.ReadFile("prompts.yaml.tmpl")
		if err != nil {
			g.logger.Errorf("error while reading the prompt config file %s", err)
			return err
		}

		// copy the default prompt file and use go template to update it
		err = os.WriteFile(promptFilePath, promptContentData, 0644)
		if err != nil {
			g.logger.Errorf("error while writing the prompt config file %s", err)
			return err
		}

		// prompt file exists, update it
		var filePaths []*string
		filePaths = append(filePaths, &promptFilePath)
		data := make(map[string]interface{})
		data["Language"] = language
		if language == "go" {
			data["GoPrompt"] = goPrompt
		} else if language == "dotnet" {
			data["DotnetPrompt"] = dotnetPrompt
		}

		err = executor.Execute(filePaths, data)
		if err != nil {
			g.logger.Errorf("error while executing the prompt file %s", err)
			return err
		}
		g.logger.Infof("Prompt file created at %s", promptFilePath)
		return nil

	}

	// g.logger.Infof("Prompt file created at %s", promptFilePath)
	return nil
}
