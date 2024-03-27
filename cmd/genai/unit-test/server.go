package unitTest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/intelops/compage/cmd/internal/utils"
)

func (u *UnitTestCmd) FetchUnitTestFromOpenAI(code string) (*models.APIResponse, error) {
	// validate the code string
	if code == "" {
		return nil, fmt.Errorf("code is empty")
	}

	// validate the language string
	if config.Language == "" {
		return nil, fmt.Errorf("language is empty")
	}

	// Fetch the `OPENAI_KEY` from the system environment
	openaiAPIKey, ok := os.LookupEnv("OPENAI_KEY")
	if !ok {
		return nil, fmt.Errorf("OPENAI_KEY is not set in the environment, please set it and try validating with `compage genaiInit` command to verify the API KEY")
	}

	// create the prompt for the OpenAI API
	prompt, err := u.fetchPromptBasedOnLanguage()
	if err != nil {
		u.logger.Errorf("error while fetching prompt: %v", err)
		return nil, err
	}
	
	// validate the prompt string
	if prompt == "" {
		return nil, fmt.Errorf("prompt is empty")
	}

	// format the prompt
	prompt = fmt.Sprintf(prompt, code)


	// create the request body
	body, err := json.Marshal(models.UnitTestRequest{
		Prompt:       prompt,
		OpenAIAPIKey: openaiAPIKey,
	})

	if err != nil {
		return nil, err
	}

	// create the request
	request, err := http.NewRequest("POST", utils.BACKEND_LLM_URL+"/unit_test_generate", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// add headers to the request
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// create the client
	client := &http.Client{}

	// send the request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error generating unit test case: %v", err)
		}
		return nil, err
	}

	var apiResponse models.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}


func (u *UnitTestCmd) fetchPromptBasedOnLanguage() (string, error) {
	if config.Language == "" {
		return "", fmt.Errorf("language is empty")
	}

	switch config.Language {
		case utils.AvailableLanguages.Go:
			return config.UnitTest.GoPrompt, nil
		case utils.AvailableLanguages.DotNet:
			return config.UnitTest.DotNetPrompt, nil
	}

	return "", fmt.Errorf("language not supported")
}
