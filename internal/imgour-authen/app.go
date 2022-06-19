package imgour_authen

import (
	"context"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/gRPC"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/auth"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/db"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/services/mail"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(env string) {
	config := Setup(env)
	restServer := route.Setup(config.Server)
	gRpcServer, gRpcLis := gRPC.Setup(config.Grpc)

	go StartGrpcServer(gRpcServer, gRpcLis)
	go StartRestfulServer(restServer)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		log.Info("Shutting down...")
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		gRpcServer.Stop()
		log.Info("Grpc server stopped")
		err := restServer.Shutdown(timeoutCtx)
		if err != nil {
			log.Errorf("Error shutting down rest server: %s\n", err)
			return
		}
		log.Info("Rest server stopped")
		log.Info("Shutdown complete")
	}
}

func StartRestfulServer(server *http.Server) {
	log.Printf("Rest server listening at %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Rest server failed: %s\n", err)
	}
}

func StartGrpcServer(s *grpc.Server, lis *net.Listener) {
	log.Printf("Grpc server listening at %v\n", (*lis).Addr())
	if err := s.Serve(*lis); err != nil {
		log.Fatalf("Grpc Server, failed to serve: %v\n", err)
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
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
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
		authConfig.ConnectionUri, authConfig.ApiKey, serverConfig.AppName, serverConfig.ServiceBaseUrl,
		serverConfig.WebsiteDomain, authConfig.ApiPath, authConfig.WebsiteBasePath, authEmailHandler, authConfig.Social)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
