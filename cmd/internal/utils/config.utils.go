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
	// check if config file exists
	if _, err := os.Stat(userHomeDir + "/.compage/gpt.env"); os.IsNotExist(err) {
		// create config file
		err = os.MkdirAll(userHomeDir+"/.compage", 0755)
		if err != nil {
			return nil, err
		}
		f, err := os.Create(userHomeDir + "/.compage/gpt.env")
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}

	vpr := viper.New()
	vpr.AddConfigPath(userHomeDir + "/.compage")
	vpr.SetConfigType("env")
	vpr.SetConfigName("gpt.env")
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
