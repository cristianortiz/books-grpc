package configs

import (
	"flag"
	"io"
	"os"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "this is config file path"
)

var (
	once         sync.Once
	cachedConfig *AppConfig
)

type ClientConfig struct {
	ClientName    string `mapstructure:"clientName"`
	LogLevel      string `mapstructure:"logLevel"`
	ServerAddress string `mapstructure:"serverAddress"`
}

type DatabaseConfig struct {
	Dbname        string         `mapstructure:"name"`
	Schema        string         `mapstructure:"schema"`
	Username      string         `mapstructure:"user"`
	Password      string         `mapstructure:"password"`
	Host          string         `mapstructure:"host"`
	Port          int            `mapstructure:"port"`
	LogMode       bool           `mapstructure:"logMode"`
	SslMode       string         `mapstructure:"sslMode"`
	Connection    ConnectionPool `mapstructure:"connectionPool"`
	MigrationPath string         `mapstructure:"migrationPath"`
}

type ConnectionPool struct {
	MaxOpenConnections int `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int `mapstructure:"maxIdleConnections"`
	MaxIdleTime        int `mapstructure:"maxIdleTime"`
	MaxLifeTime        int `mapstructure:"maxLifeTime"`
	TimeOut            int `mapstructure:"timeout"`
}

type ServerConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	LogLevel    string `mapstructure:"logLevel"`
}
type AppConfig struct {
	ServerConfig ServerConfig   `mapstructure:"app"`
	DBConfig     DatabaseConfig `mapstructure:"db"`
	ClientConfig ClientConfig   `mapstructure:"client"`
}

func ProvideAppConfig() (c *AppConfig, err error) {
	once.Do(func() {
		var configFile string
		flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
		flag.Parse()

		var configReader io.ReadCloser
		configReader, err = os.Open(configFile)
		defer configReader.Close() //nolint:staticcheck
		if err != nil {
			return
		}
		c, err = LoadConfig(configReader)
		if err != nil {
			return
		}
		cachedConfig = c

	})
	return cachedConfig, err
}

func LoadConfig(reader io.Reader) (*AppConfig, error) {
	var appconfig AppConfig
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	keysToEnvironmentVariables := map[string]string{
		"app.port":    "APP_PORT",
		"db.name":     "DB_NAME",
		"db.user":     "DB_USER",
		"db.host":     "DB_HOST",
		"db.port":     "DB_PORT",
		"db.password": "DB_PASSWORD",
	}
	err := bind(keysToEnvironmentVariables)
	if err != nil {
		return nil, err
	}

	if err := viper.ReadConfig(reader); err != nil {
		return nil, errors.Wrap(err, "failed to load app config file")
	}
	if err := viper.Unmarshal(&appconfig); err != nil {
		return nil, errors.Wrap(err, "unable to parse app config file")
	}
	return &appconfig, nil
}

func bind(keysToEnvironmentVariables map[string]string) error {
	var bindErrors error

	for key, environmentVar := range keysToEnvironmentVariables {
		if err := viper.BindEnv(key, environmentVar); err != nil {
			bindErrors = multierror.Append(bindErrors, err)
		}
	}
	return bindErrors
}
