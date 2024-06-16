package config

import "github.com/spf13/viper"

func NewViper() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
