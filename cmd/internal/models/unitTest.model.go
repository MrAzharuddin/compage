package models

type Data struct {
	Code string `json:"code"`
}

type FileStructure struct {
	FileName string              `json:"fileName"`
	Data     map[string][]string `json:"data"`
}

type APIResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type UnitTestRequest struct {
	Prompt       string `json:"prompt"`
	OpenAIAPIKey string `json:"openai_api_key"`
}
