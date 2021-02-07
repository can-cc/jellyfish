package configs

import (
	"strings"

	"github.com/juju/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Application  ApplicationConfig
	DataSource   DataSourceConfig
	Logger       LoggerConfig
	Storage      StorageConfig
	Bucket       BucketConfig
	Notification NotificationConfig
}

type DataSourceConfig struct {
	RDS RDSConfig `mapstructure:"rds"`
}

type RDSConfig struct {
	DatabaseUrl string `mapstructure:"database_url"`
}

type ApplicationConfig struct {
	Addr         string
	JwtSecret    string `mapstructure:"jwt_secret"`
	JwtHeaderKey string `mapstructure:"jwt_header_key"`
}

type LoggerConfig struct {
	Level       string
	OutputPaths []string `mapstructure:"output_paths"`
}

type StorageConfig struct {
	Endpoint          string `mapstructure:"endpoint"`
	AccessKeyID       string `mapstructure:"access_key_id"`
	SecretAccessKeyID string `mapstructure:"secret_access_key_id"`
	UseSSL            bool   `mapstructure:"use_ssl"`
	Location          string `mapstructure:"location"`
}

type BucketConfig struct {
	Image string
}

type NotificationConfig struct {
	Endpoint string
}

func loadConfig(file string) (*AppConfig, error) {
	var config AppConfig

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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
