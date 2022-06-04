package imgour_authen

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
	log "github.com/sirupsen/logrus"
	"os"
)

func Start(env string) {
	Setup(env)
	r := route.Setup()

	err := r.Run(":8018")
	if err != nil {
		log.Error(err)
	}
}

func Setup(env string) {
	config := LoadConfig(env)
	logInit()
	dbSetup(config.Db.Host, config.Db.Username, config.Db.Password)
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
