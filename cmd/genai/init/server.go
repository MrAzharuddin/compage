package init

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/intelops/compage/cmd/internal/utils"
)

func (ga *GenAIStart) CheckOpenAITokens() error {
	if OPENAI_KEY == "" {
		ga.logger.Error("OPENAI_KEY is not set")
		return errors.New("OPENAI_KEY is not set")
	}

	// create a request body to send to the server
	tokenbody := models.TokenSchema{
		OpenAIApiKey: OPENAI_KEY,
	}

	// convert the request body to JSON
	tokenbodyJSON, err := json.Marshal(tokenbody)
	if err != nil {
		ga.logger.Error(err)
		return err
	}

	response, err := http.Post(utils.BACKEND_LLM_URL+"/validate_openai", "application/json", bytes.NewBuffer(tokenbodyJSON))
	if err != nil {
		ga.logger.Error(err)
		return err
	}

	defer response.Body.Close()

	// read the response status code
	statusCode := response.StatusCode

	// read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		ga.logger.Error(err)
		return err
	}

	// if the status code is not 200, throw an error
	if statusCode != 200 {
		ga.logger.Error(string(body))
		return errors.New(string(body))
	}

	ga.logger.Info("📢 OpenAI API key validated successfully! Please explore our commands to get started. ⚡🚀")

	return nil
}
