package repository

import (
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type etaxTableRepositoryDb struct {
	db *gorm.DB
}

func NewEtaxTableRepositoryDb(db *gorm.DB) etaxTableRepositoryDb {
	return etaxTableRepositoryDb{db: db}
}

func (rp etaxTableRepositoryDb) SqlGetAll() ([]EtaxTable, error) {
	etaxTables := []EtaxTable{}
	r := rp.db.Table("EA_TMPINVOICEETAXTABLE").Find(&etaxTables, "STATUS_SIGN = ?", 0)
	if r.Error != nil {
		return nil, r.Error
	}
	return etaxTables, nil
}

func (rp etaxTableRepositoryDb) SqlUpdate(etaxTable *EtaxTable) error {
	r := rp.db.Table("EA_TMPINVOICEETAXTABLE").Model(&etaxTable).Update("STATUS_SIGN", 1)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
