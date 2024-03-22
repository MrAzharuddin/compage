package utils

import (
	"os"

	"github.com/intelops/compage/cmd/internal/models"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	Viper *viper.Viper
}


func LoadViper() (*ViperConfig, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := userHomeDir + "/.compage"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return nil, err
	}

	vpr := viper.New()
	vpr.AddConfigPath(userHomeDir + "/.compage")
	vpr.SetConfigType("yaml")
	vpr.SetConfigName("prompts.yaml")
	vpr.AutomaticEnv()
	if err := vpr.ReadInConfig(); err != nil {
		return nil, err
	}
	return &ViperConfig{
		Viper: vpr,
	}, nil
}

func (vc *ViperConfig) GetConfig() *ViperConfig {
	return vc
}

// write a variable in config path
func (vc *ViperConfig) WriteConfig(key string, value string) error {
	vc.Viper.Set(key, value)
	return vc.Viper.WriteConfig()
}

// read a variable from config path
func (vc *ViperConfig) ReadConfig(key string) (string, error) {
	return vc.Viper.GetString(key), nil
}

// unmarshal all the configs from file
func (vc *ViperConfig) Unmarshal() (config *models.EnvConfig, err error) {
	err = vc.Viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return
}
