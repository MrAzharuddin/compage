package models

type EnvConfig struct {
	OpenAIAccessToken string `mapstructure:"OPENAI_ACCESS_TOKEN"`
}

func (ec *EnvConfig) Validate() bool {
	return ec.OpenAIAccessToken != ""
}
