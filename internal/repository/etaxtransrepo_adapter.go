package repository

import (
	"go-etax/internal/logs"

	"gorm.io/gorm"
)

type etaxTransRepositoryDb struct {
	db *gorm.DB
}

func NewEtaxTransRepositoryDb(db *gorm.DB) *etaxTransRepositoryDb {
	return &etaxTransRepositoryDb{db: db}
}

func (rp *etaxTransRepositoryDb) GetById(document_id string, company string) ([]EtaxTransRepositoryDb, error) {
	i := []EtaxTransRepositoryDb{}
	tx := rp.db.Table("EA_TMPINVOICEETAXTRANS").Find(&i, "DOCUMENT_ID = ? and COMPANY = ?", document_id, company)
	if tx.Error != nil {
		return nil, tx.Error
	}
	logs.Info("SELECT ETAX_TRANS ON " + document_id)
	return i, nil
}
