package docgeneration

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

// docGenerationServer sends a POST request to the backend server to generate documentation
// for a given folderStructure. It expects a JSON response with the generated documentation
// in the 'result' field
func (dg *DocGenerationCmd) docGenerationServer(folderStructure string) (*models.DocGenResponse, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Validate the folderStructure string
	if folderStructure == "" {
		return nil, fmt.Errorf("folderStructure is empty")
	}

	// Fetch the `OPENAI_KEY` from the system environment
	openaiAPIKey, ok := os.LookupEnv("OPENAI_KEY")
	if !ok {
		return nil, fmt.Errorf("OPENAI_KEY is not set in the environment, please set it and try validating with `compage genaiInit` command to verify the API KEY")
	}

	// Fetch Prompt from json file
	prompt, err := dg.fetchPrompt("prompts.json")
	if err != nil {
		return nil, err
	}

	if prompt == "" {
		return nil, fmt.Errorf("prompt is empty")
	}

	// Format the prompt by replacing the placeholder with the folderStructure string
	prompt = fmt.Sprintf(prompt, folderStructure)

	// Create the request body
	body, err := json.Marshal(models.UnitTestRequest{
		Prompt:       prompt,
		OpenAIAPIKey: openaiAPIKey,
	})

	if err != nil {
		return nil, err
	}

	// Create the request
	request, err := http.NewRequest("POST", utils.BACKEND_LLM_URL+"/doc_generate", bytes.NewBuffer(body))
	if err != nil {
		dg.logger.Error("Error creating request:", err)
		return nil, err
	}

	// Set the request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		dg.logger.Error("Error sending request:", err)
		return nil, err
	}

	defer response.Body.Close()

	// Read the response body
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

func (dg *DocGenerationCmd) fetchPrompt(filepath string) (string, error) {
	prompts, err := models.ReadPromptsConfigJSONFile(filepath)
	if err != nil {
		return "", err
	}
	return prompts.Documentation.DocPrompt, nil
}
