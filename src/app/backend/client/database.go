package client

import (
	"fmt"
	"log"

	"github.com/Triticumdico/dashboard/src/app/backend/args"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ClientDb = &clientDb{}

type clientDb struct {
	clientDb *gorm.DB
}

func (self *clientDb) OpenDbConnection(dbName string) *gorm.DB {

	var err error

	databaseurl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		args.Config.GetHostDatabase(),
		args.Config.GetUsernameDatabase(),
		args.Config.GetPasswordDatabase(),
		dbName,
		args.Config.GetPortDatabase())

	self.clientDb, err = gorm.Open(postgres.Open(databaseurl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error while connecting database ", err)
	} else {
		log.Printf("Connection establish to database")
	}

	return self.clientDb
}

func NewClientDb() *clientDb {
	return ClientDb
}
