package models

type Data struct {
	Code string `json:"code"`
}
type APIResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type UnitTestRequest struct {
	Prompt   string `json:"prompt"`
	Language string `json:"language"`
}
