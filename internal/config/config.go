package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("cannot read config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("cannot unmarshal config: %w", err)
	}

	return config, nil
}
