package unitTest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/intelops/compage/cmd/internal/utils"
)

func (u *UnitTestCmd) FetchUnitTestFromOpenAI(code string) (*models.APIResponse, error) {
	client := &http.Client{}
	// validate the code string
	if code == "" {
		return nil, errors.New("code is empty")
	}

	// Fetch the `OPENAI_KEY` from the system environment
	openaiAPIKey, ok := os.LookupEnv("OPENAI_KEY")
	if !ok {
		return nil, errors.New("OPENAI_KEY is not set in the environment, please set it and try validating with `compage genaiInit` command to verify the API KEY")
	}

	// create the prompt for the OpenAI API
	prompt := fmt.Sprintf("Write a unit test case for the following Golang programming language code using the in-built testing package in golang:\n%s. Make sure the unit test case you are generating is providing the imports on the top and also keep that whole test case in between three backticks(```) at the beginning and end of the unit test case.", code)

	// create the request body
	body, err := json.Marshal(models.UnitTestRequest{
		Prompt:   prompt,
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
