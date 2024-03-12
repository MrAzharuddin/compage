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
	if OPENAI_KEY == "" || USERNAME == "" {
		return errors.New("OPENAI_KEY or USERNAME is not set")
	}

	// create a request body to send to the server
	tokenbody := models.TokenSchema{
		Username: USERNAME,
		OpenAIKey: OPENAI_KEY,
	}

	// convert the request body to JSON
	tokenbodyJSON, err := json.Marshal(tokenbody)
	if err != nil {
		return err
	}

	response, err := http.Post(utils.BACKEND_LLM_URL+"/create-token", "application/json", bytes.NewBuffer(tokenbodyJSON))
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read the response status code
	statusCode := response.StatusCode

	// read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// if the status code is not 200, throw an error
	if statusCode != 200 {
		return errors.New(string(body))
	}

	// if the status code is 200, return nil
	var resdata models.TokenResponseSchema
	err = json.Unmarshal(body, &resdata)
	if err != nil {
		return err
	}

	err = ga.viperConfig.WriteConfig("OPENAI_ACCESS_TOKEN", resdata.AccessToken )
	if err != nil {
		return err
	}

	ga.logger.Info("Tokens created successfully")
	
	return nil
}
