package db

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	log "github.com/sirupsen/logrus"
)

var (
	client *qmgo.Client
	userDb *qmgo.Collection
)

func Setup(host, username, password string) {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, host)
	var err error
	client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: uri})
	userDb = client.Database("imgour-authen").Collection("users")
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
