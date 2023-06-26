package main

import (
	"fmt"
	"github.com/ariashabry/KreditPlus/handlers"
	"github.com/ariashabry/KreditPlus/helpers/env"
	"github.com/ariashabry/KreditPlus/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	//DB Connection
	cs := env.Get().ConnectionString()
	db, err := gorm.Open(mysql.Open(cs), &gorm.Config{})

	//check DB Connection
	if err != nil {
		log.Fatal("Connection To DB Failed")
		return
	}
	log.Println("DB Connected Successfully")

	//migration models to db
	err = models.MigrateModel(db)

	if err != nil {
		log.Fatal("Migration Model Failed")
		return
	}
	log.Println("Migration DB Successfully")

	//set echo framework to routing
	e := echo.New()

	//set handler or routing
	h := handlers.Context{Echo: e, DB: db}
	h.Api("api")

	//set address
	Host := env.Get().AppHost
	Port := env.Get().AppPort
	Address := fmt.Sprintf("%s:%d", Host, Port)

	log.Printf("Start listen and serve at %s: %v", Host, Port)
	err = e.Start(Address)
	if err != nil {
		log.Fatal("failed to connect to server")
		return
	}
}
