package main

import (
	"fmt"
	"go-etax/repository"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", "sa", "P@ssw0rd", "127.0.0.1", 1433, "TestDB")
	db, err := gorm.Open(sqlserver.Open(dsn))
	if err != nil {
		panic(err)
	}

	etaxTableRepository := repository.NewEtaxTableRepositoryDb(db)
	_ = etaxTableRepository

	etaxTables, err := etaxTableRepository.SqlGetAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(etaxTables)

	etaxTableRepository.SqlUpdate(&etaxTables[0])
}
