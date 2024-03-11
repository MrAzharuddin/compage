package testcmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/intelops/compage/cmd/models"
)

const llmBaseURL = "http://localhost:8000/api"

func Server(accessToken string, code string) (models.CodeGenResponse, error) {
	client := &http.Client{}
	if code == "" {
		return models.CodeGenResponse{}, fmt.Errorf("code is empty")
	}

	prompt := fmt.Sprintf("Write a unit test case for the following Golang programming language code using the in-built testing package in golang:\n%s. Make sure the unit test case you are generating is providing the imports on the top and also keep that whole test case in between three backticks(```) at the beginning and end of the unit test case.", code)

	requestBody, err := json.Marshal(map[string]string{
		"prompt":   prompt,
		"language": "go",
	})

	if err != nil {
		return models.CodeGenResponse{}, err
	}

	req, err := http.NewRequest("POST", llmBaseURL+"/code_generate", bytes.NewBuffer(requestBody))
	if err != nil {
		return models.CodeGenResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return models.CodeGenResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if resp.StatusCode != 200 {
			return models.CodeGenResponse{}, fmt.Errorf("error generating unit test: %v", err)
		}
		return models.CodeGenResponse{}, err
	}

	var response models.CodeGenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.CodeGenResponse{}, err
	}

	return response, nil
}
