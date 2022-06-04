package imgour_authen

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Db DbConfig `mapstructure:"db"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func LoadConfig(env string) *Config {
	configName := env + ".config"
	viper.SetConfigName(configName)       // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath("./configs")      // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return &config
}
