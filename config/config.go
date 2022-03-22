package config

import (
	"os"

	"github.com/spf13/viper"
)

//Config ...
type Config struct {
	Token string `mapstructure:"FEATWS_RESOLVER_CLIMATEMPO_TOKEN"`
	Port  string `mapstructure:"PORT"`
}

//LoadConfig ...
func LoadConfig(config *Config) (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetDefault("FEATWS_RESOLVER_CLIMATEMPO_TOKEN", "")
	viper.SetDefault("PORT", "7000")

	err = viper.ReadInConfig()
	if err != nil {
		if err2, ok := err.(*os.PathError); !ok {
			err = err2
			return
		}
	}

	err = viper.Unmarshal(config)

	return
}
