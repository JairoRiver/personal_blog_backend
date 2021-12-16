package util

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"'`
}

func LoadConfig(path, name string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
