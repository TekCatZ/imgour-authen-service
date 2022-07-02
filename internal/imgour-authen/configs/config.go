package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
	Auth   AuthConfig   `mapstructure:"auth"`
	Mail   MailConfig   `mapstructure:"mail"`
	Grpc   GrpcConfig   `mapstructure:"grpc-server"`
}

type CollectionsName struct {
	Users string `mapstructure:"users"`
}

type DbConfig struct {
	Host       string          `mapstructure:"host"`
	Username   string          `mapstructure:"username"`
	Password   string          `mapstructure:"password"`
	DbName     string          `mapstructure:"dbname"`
	Collection CollectionsName `mapstructure:"collections_name"`
}

type ServerConfig struct {
	Port           string `mapstructure:"port"`
	AppName        string `mapstructure:"app_name"`
	ServiceName    string `mapstructure:"service_name"`
	WebsiteDomain  string `mapstructure:"website_domain"`
	ServiceBaseUrl string `mapstructure:"service_base_url"`
}

type AuthConfig struct {
	ConnectionUri   string        `mapstructure:"connection_uri"`
	ApiKey          string        `mapstructure:"api_key"`
	ApiPath         string        `mapstructure:"api_path"`
	WebsiteBasePath string        `mapstructure:"website_base_path"`
	Social          SocialConfigs `mapstructure:"social"`
}

type MailConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Server   string `mapstructure:"server"`
	Port     string `mapstructure:"port"`
}

type SocialConfigs struct {
	GoogleConfigs map[string]SocialConfig `mapstructure:"google"`
	GithubConfigs map[string]SocialConfig `mapstructure:"github"`
}

type SocialConfig struct {
	Id     string `mapstructure:"id"`
	Secret string `mapstructure:"secret"`
}

type GrpcConfig struct {
	Port string `mapstructure:"port"`
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
