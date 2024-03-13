package unitTest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/intelops/compage/cmd/internal/utils"
)

func (u *UnitTestCmd) FetchUnitTestFromOpenAI(code string) (*models.APIResponse, error) {
	client := &http.Client{}
	// validate the code string
	if code == "" {
		return nil, errors.New("code is empty")
	}

	// get the access token from the config
	config, err := u.viperConfig.Unmarshal()
	if err != nil {
		return nil, err
	}
	// store the access token from the config
	accessToken := config.OpenAIAccessToken

	// validate the access token
	if accessToken == "" {
		return nil, errors.New("access token is empty")
	}

	// create the prompt for the OpenAI API
	prompt := fmt.Sprintf("Write a unit test case for the following Golang programming language code using the in-built testing package in golang:\n%s. Make sure the unit test case you are generating is providing the imports on the top and also keep that whole test case in between three backticks(```) at the beginning and end of the unit test case.", code)

	// create the request body
	body, err := json.Marshal(models.UnitTestRequest{
		Prompt:   prompt,
		Language: "go",
	})

	if err != nil {
		return nil, err
	}

	// create the request
	request, err := http.NewRequest("POST", utils.BACKEND_LLM_URL+"/code_generate", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// add headers to the request
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
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
