package models

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type PromptsConfig struct {
	Language      string              `json:"language"`
	UnitTest      UnitTestConfig      `json:"unitTest"`
	Documentation DocumentationConfig `json:"documentation"`
}

type UnitTestConfig struct {
	GoPrompt     string `json:"goPrompt"`
	DotNetPrompt string `json:"dotnetPrompt"`
}

type DocumentationConfig struct {
	DocPrompt string `json:"docPrompt"`
}

func ReadPromptsConfigJSONFile(configFile string) (*PromptsConfig, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Errorf("error reading config file: %v", err)
		return nil, err
	}
	var config *PromptsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Errorf("error unmarshalling config file: %v", err)
		return nil, err
	}
	return config, nil
}
