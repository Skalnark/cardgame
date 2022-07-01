package main

import (
	"fmt"
	database "skalnark/cards/providers"
	"skalnark/cards/web"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	DB, err := gorm.Open(sqlite.Open("database/card_game.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't open the database: ", err)
		return
	}

	database.DbMigrate(DB)

	host := "localhost:3000"

	web.RunServer(DB, host)

}
