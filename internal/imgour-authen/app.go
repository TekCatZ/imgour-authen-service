package imgour_authen

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func Start() {
	r := route.Setup()

	err := r.Run(":8018")
	if err != nil {
		log.Error(err)
	}
}
