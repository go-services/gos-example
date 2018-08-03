package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"github.com/go-services/gos-project/services/user/config"
)

var session *gorm.DB

func Session() *gorm.DB {
	if session != nil {
		return session
	}
	s, err := Connect()
	if err != nil {
		log.Fatalf("could not connect to the DB: %v", err)
	}
	session = s
	return session
}

func Connect() (*gorm.DB, error) {
	return gorm.Open("mysql", config.Get().DBConnectionString)
}

func Close() error {
	if session != nil {
		err := session.Close()
		if err != nil {
			session = nil
			return err
		}
		session = nil
	}
	return nil
}
