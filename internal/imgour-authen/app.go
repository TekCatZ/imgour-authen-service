package imgour_authen

import (
	"fmt"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/auth"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/mail"
	log "github.com/sirupsen/logrus"
	"os"
)

func Start(env string) {
	config := Setup(env)
	r := route.Setup()

	err := r.Run(fmt.Sprintf(":%s", config.Server.Port))
	if err != nil {
		log.Error(err)
	}
}

func Setup(env string) *configs.Config {
	config := configs.LoadConfig(env)
	logInit()
	dbSetup(config.Db.Host, config.Db.Username, config.Db.Password)
	mailSetup(config.Mail)
	authSetup(config.Auth, config.Server)

	return config
}

func logInit() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func dbSetup(host, username, password string) {
	db.Setup(host, username, password)
	db.Ping()
}

func mailSetup(mailConfig configs.MailConfig) {
	err := mail.Setup(mailConfig)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func authSetup(authConfig configs.AuthConfig, serverConfig configs.ServerConfig) {
	authEmailHandler := mail.GetAuthEmailHandler()

	err := auth.Setup(
		authConfig.ConnectionUri, authConfig.ApiKey, serverConfig.AppName, serverConfig.ApiDomain,
		serverConfig.WebsiteDomain, authConfig.ApiPath, authConfig.WebsiteBasePath, authEmailHandler, authConfig.Social)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
