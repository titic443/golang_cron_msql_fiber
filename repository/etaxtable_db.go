package repository

import (
	"log"

	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type etaxTableRepositoryDb struct {
	db *gorm.DB
}

func NewEtaxTableRepositoryDb(db *gorm.DB) etaxTableRepositoryDb {
	return etaxTableRepositoryDb{db: db}
}

func (p etaxTableRepositoryDb) SqlGetAll() ([]EtaxTable, error) {
	etaxTables := []EtaxTable{}
	r := p.db.Table("EA_TMPINVOICEETAXTABLE").Find(&etaxTables, "STATUS_SIGN = ?", 0)
	if r.Error != nil {
		log.Println(r.Error)
	}
	return etaxTables, nil
}

// func (p etaxTableRepositoryDb) SqlGetById(id int) (*EtaxTable, error) {
// 	return nil, nil
// }

func (p etaxTableRepositoryDb) SqlUpdate(etaxTable *EtaxTable) error {
	r := p.db.Table("EA_TMPINVOICEETAXTABLE").Model(&etaxTable).Update("STATUS_SIGN", 1)
	if r.Error != nil {
		log.Println(r.Error)
		return r.Error
	}
	return nil
}
