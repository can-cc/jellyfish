package configs

import (
	"github.com/juju/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Application ApplicationConfig
	DataSource  DataSourceConfig
	Logger      LoggerConfig
}

type DataSourceConfig struct {
	RDS RDSConfig `mapstructure:"rds"`
}

type RDSConfig struct {
	DatabaseUrl string `mapstructure:"database_url"`
}

type ApplicationConfig struct {
	Addr      string
	JwtSecret string
}

type LoggerConfig struct {
	Level            string
	ErrorOutputPaths []string `mapstructure:"error_output_paths"`
}

func loadConfig(file string) (*AppConfig, error) {
	var config AppConfig

	viper.SetEnvPrefix("jfish")
	viper.AutomaticEnv()
	viper.SetConfigFile(file)

	err := viper.ReadInConfig()

	if err != nil {
		return nil, errors.Trace(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &config, nil
}

func LoadConfig(file string) (*AppConfig, error) {
	return loadConfig(file)
}
