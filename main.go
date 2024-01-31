package main

import (
	"fmt"
	"go-etax/handler"
	"go-etax/repository"
	"go-etax/service"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hirochachacha/go-smb2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	conn, err := net.Dial("tcp", "10.15.5.4:445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "titi.cha",
			Password: "For+ever16!",
			Domain:   "energyabsolute",
		},
	}

	client, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer c.Logoff()

	fileshareRepository := repository.NewfileshareRepository(client)
	// fileshareRepository.DownloadFile("smb://10.15.5.4/it-data$/TESTApp")
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	_, err = os.ReadDir("./download")
	if err != nil {
		os.Mkdir("./download", 0777)
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", "sa", "P@ssw0rd", "127.0.0.1", 1433, "TestDB")
	db, err := gorm.Open(sqlserver.Open(dsn))
	if err != nil {
		panic(err)
	}

	etaxTableRepository := repository.NewEtaxTableRepositoryDb(db)
	etaxTransRepository := repository.NewEtaxTransRepositoryDb(db)
	etaxTableService := service.NewEtaxTableService(etaxTableRepository, etaxTransRepository)
	etaxTableHandler := handler.NewEtaxTableHandler(etaxTableService)
	app.Get("/etax", etaxTableHandler.SendEtaxToEco)

	app.Listen(":8888")
}
