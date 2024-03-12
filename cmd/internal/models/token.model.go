package models

type TokenSchema struct {
	Username  string `json:"username"`
	OpenAIKey string `json:"api_key"`
}

type TokenResponseSchema struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}
