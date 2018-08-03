package main

import (
	"time"
	serviceDB "github.com/go-services/gos-project/services/user/db"
	"fmt"
	"github.com/go-services/gos-project/services/user/config"
	"github.com/go-services/gos-project/services/user/cmd/migration/db"
)

func main() {
	fmt.Println(config.Get().DBConnectionString)
	for _, err := serviceDB.Connect(); err != nil; {
		fmt.Println(fmt.Sprintf("Could not connect to the DB trying again in 10 sec: %v", err))
		time.Sleep(10 * time.Second)
	}
	if config.Get().ENV == "dev" {
		fmt.Println("Droping DB")
		db.DropAll()
		fmt.Println("Migrating DB")
		db.Migrate()
		fmt.Println("Seeding DB")
		db.Seed()
	} else {
		fmt.Println("Migrating Production DB")
		db.Migrate()
		fmt.Println("Seeding Production DB")
		db.SeedProd()
	}
}
