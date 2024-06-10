package repository

import (
	"go-etax/internal/logs"

	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type etaxTableRepositoryDb struct {
	db *gorm.DB
}

func NewEtaxTableRepositoryDb(db *gorm.DB) *etaxTableRepositoryDb {
	return &etaxTableRepositoryDb{db: db}
}

func (rp *etaxTableRepositoryDb) SqlGetAll() ([]EtaxTable, error) {
	etaxTables := []EtaxTable{}
	r := rp.db.Table("TMPINVOICEETAXTABLE").Find(&etaxTables, "STATUS_SIGN = ?", 0)
	if r.Error != nil {
		return nil, r.Error
	}
	logs.Info("SELECT STATUS_SIGN = 0")
	return etaxTables, nil
}

func (rp *etaxTableRepositoryDb) SqlUpdate(docId string) error {
	r := rp.db.Table("TMPINVOICEETAXTABLE").Where("DOCUMENT_ID = ?", docId).Update("STATUS_SIGN", 1)
	// r := rp.db.Table("EA_TMPINVOICEETAXTABLE").Model(&etaxTable).Update("STATUS_SIGN", 1)
	if r.Error != nil {
		return r.Error
	}
	logs.Info("Update STATUS_SIGN = 1 ON " + docId)
	return nil
}
