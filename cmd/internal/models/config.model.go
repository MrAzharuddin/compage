package models

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type GoUnitTestConfig struct {
	GoPrompt string `json:"goPrompt"`
}

type GoDocumentationConfig struct {
	GoDocPrompt string `json:"goDocPrompt"`
}

type GoPromptsConfig struct {
	UnitTest      GoUnitTestConfig      `json:"unitTest"`
	Documentation GoDocumentationConfig `json:"documentation"`
}

// Dotnet Models
type DotnetUnitTestConfig struct {
	DotnetPrompt string `json:"dotnetPrompt"`
}

type DotnetDocumentationConfig struct {
	DotnetDocPrompt string `json:"dotnetDocPrompt"`
}

type DotnetPromptsConfig struct {
	UnitTest      DotnetUnitTestConfig      `json:"unitTest"`
	Documentation DotnetDocumentationConfig `json:"documentation"`
}

func ReadGoConfigJSONFile(configFile string, language string) (*GoPromptsConfig, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Errorf("error reading config file: %v", err)
		return nil, err
	}
	var config *GoPromptsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Errorf("error unmarshalling config file: %v", err)
		return nil, err
	}
	return config, nil
}

func ReadDotnetConfigJSONFile(configFile string, language string) (*DotnetPromptsConfig, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Errorf("error reading config file: %v", err)
		return nil, err
	}
	var config *DotnetPromptsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Errorf("error unmarshalling config file: %v", err)
		return nil, err
	}
	return config, nil
}
