package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or enviroment variable.
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	HostName             string        `mapstructure:"HOST_NAME"`
	AwsRegion            string        `mapstructure:"AWS_REGION"`
	AwsKey               string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecret            string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsBucket            string        `mapstructure:"AWS_BUCKET_NAME"`
}

// LoadConfig reads configuration from file or envioroment variables.
func LoadConfig(path, name string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
