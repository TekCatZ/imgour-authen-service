package db

import (
	"context"
	"fmt"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/qiniu/qmgo"
	log "github.com/sirupsen/logrus"
)

var (
	client *qmgo.Client
	userDb *qmgo.Collection
)

func Setup(dbConfig configs.DbConfig) {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbConfig.Username, dbConfig.Password, dbConfig.Host)
	var err error
	client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})
	userDb = client.Database(dbConfig.DbName).Collection(dbConfig.Collection.Users)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

func Ping() bool {
	err := client.Ping(10)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}
