package docgeneration

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

func (dg *DocGenerationCmd) docGenerationServer(folderStructure string) (*models.DocGenResponse, error) {
	client := &http.Client{}

	// validate the folderStructure string
	if folderStructure == "" {
		return nil, errors.New("folderStructure is empty")
	}

	// Fetch the `OPENAI_KEY` from the system environment
	openaiAPIKey, ok := os.LookupEnv("OPENAI_KEY")
	if !ok {
		return nil, errors.New("OPENAI_KEY is not set in the environment, please set it and try validating with `compage genaiInit` command to verify the API KEY")
	}

	// create the prompt for the OpenAI API
	prompt := fmt.Sprintf("Generate documentation for the following folder structure and also provide code flow diagram in the mermaid format. The entire documentation should be in markdown format. \n\nFolder Structure : \n%s", folderStructure)

	// create the request body
	body, err := json.Marshal(models.UnitTestRequest{
		Prompt: prompt,
		OpenAIAPIKey: openaiAPIKey,
	})

	if err != nil {
		return nil, err
	}

	// create the request
	request, err := http.NewRequest("POST", utils.BACKEND_LLM_URL+"/doc_generate", bytes.NewBuffer(body))
	if err != nil {
		dg.logger.Error("Error creating request:", err)
		return nil, err
	}

	// set the request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		dg.logger.Error("Error sending request:", err)
		return nil, err
	}

	defer response.Body.Close()

	// read the response body
	body, err = io.ReadAll(response.Body)

	if err != nil {
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error generating documentation: %v", err)
		}
		return nil, err
	}

	var apiResponse models.DocGenResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
