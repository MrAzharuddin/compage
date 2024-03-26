package models

type DocGenData struct {
	Docs string `json:"docs"`
}

type DocGenResponse struct {
	Status  string     `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DocGenData `json:"data"`
}
