package main

import (
	"fmt"
	"go-etax/handler"
	"go-etax/repository"
	"go-etax/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", "sa", "P@ssw0rd", "127.0.0.1", 1433, "TestDB")
	db, err := gorm.Open(sqlserver.Open(dsn))
	if err != nil {
		panic(err)
	}

	etaxTableRepository := repository.NewEtaxTableRepositoryDb(db)
	etaxTableService := service.NewEtaxTableService(etaxTableRepository)
	etaxTableHandler := handler.NewEtaxTableHandler(etaxTableService)

	app.Get("/etax", etaxTableHandler.SendEtaxToEco)

	app.Listen(":8888")
}
