package main

import (
	"fmt"
	"go-etax/handler"
	"go-etax/internal/logs"
	"go-etax/internal/repository"
	"go-etax/internal/service"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	var err error
	initConfig()
	initTimeZone()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	_, err = os.ReadDir("./download")
	if err != nil {
		os.Mkdir("./download", 0777)
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s&encrypt=disable&connection+timeout=30", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.hostname"), viper.GetInt("db.port"), viper.GetString("db.db"))
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	etaxTableRepository := repository.NewEtaxTableRepositoryDb(db)
	etaxTransRepository := repository.NewEtaxTransRepositoryDb(db)
	etaxTableService := service.NewEtaxTableService(etaxTableRepository, etaxTransRepository)
	etaxTableHandler := handler.NewEtaxTableHandler(etaxTableService)
	app.Get("/etax", etaxTableHandler.SendEtaxToEco)

	logs.Info("App Sign ETAX listening on port" + viper.GetString("app.port"))
	logs.Info("Create cronjob (0 */30 * * *)")
	c := cron.New()
	c.AddFunc("0 */30 * * *", func() {
		etaxTableHandler.SendEtaxToEcoCronjob()
		logs.Info("[Job 1]Every 30 minute job\n")
	})

	logs.Info("Start cronjob (0 */30 * * *)")

	c.Start()
	app.Listen(":8888")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
