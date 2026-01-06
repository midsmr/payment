package application

import (
	"strings"

	"github.com/spf13/viper"
)

type AllConfig struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Name  string `mapstructure:"name"`
	Debug bool   `mapstructure:"debug"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

var Config AllConfig

func initConfig() error {
	v := viper.New()

	v.SetEnvPrefix("PAYMENT")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(Workdir)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	return v.Unmarshal(&Config)
}
