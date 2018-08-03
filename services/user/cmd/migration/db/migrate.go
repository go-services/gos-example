package db

import (
	"github.com/go-services/gos-project/services/user/db"
	"github.com/go-services/gos-project/services/user/db/models"
	"log"
)

func Migrate() {
	defer db.Close()
	transaction := db.Session().AutoMigrate(&models.User{})
	if transaction.Error != nil {
		log.Fatal(transaction.Error)
	}
	ForeignKeys()
}

func ForeignKeys() {

}
func DropAll() {
	defer db.Close()
	transaction := db.Session().DropTableIfExists(
		&models.User{},
	)
	if transaction.Error != nil {
		log.Fatal(transaction.Error)
	}
}