package repository

type EtaxTransRepositoryDb struct {
	COMPANY            string  `gorm:"column:COMPANY;type:nvarchar(10);default:'';not null" json:"company"`
	DOCUMENT_ID        string  `gorm:"column:DOCUMENT_ID;type:nvarchar(20);default:'';not null" json:"document_id" `
	LINE_BASE_AMOUNT   float64 `gorm:"column:LINE_BASE_AMOUNT;type:numeric(32,16);default:0;not null"  json:"line_base_amount"`
	LINE_TAX_AMOUNT    float64 `gorm:"column:LINE_TAX_AMOUNT;type:numeric(32,16);default:0;not null" json:"line_tax_amount" `
	LINE_TAX_RATE      float64 `gorm:"column:LINE_TAX_RATE;type:numeric(32,16);default:0;not null"  json:"line_tax_rate"`
	LINE_TAX_TYPE_CODE string  `gorm:"column:LINE_TAX_TYPE_CODE;type:nvarchar(10);default:'';not null"  json:"line_tax_type_code"`
	LINE_TOTAL_AMOUNT  float64 `gorm:"column:LINE_TOTAL_AMOUNT;type:numeric(32,16);default:0;not null"  json:"line_total_amount"`
	PRODUCT_CODE       string  `gorm:"column:PRODUCT_CODE;type:nvarchar(20);default:'';not null" json:"product_code" `
	PRODUCT_NAME       string  `gorm:"column:PRODUCT_NAME;type:nvarchar(90);default:'';not null"  json:"product_name"`
	PRODUCT_PRICE      float64 `gorm:"column:PRODUCT_PRICE;type:numeric(32,16);default:0;not null"  json:"product_price"`
	PRODUCT_QUANTITY   float64 `gorm:"column:PRODUCT_QUANTITY;type:numeric(32,16);default:0;not null"  json:"product_quantity"`
	RECVERSION         int     `gorm:"column:RECVERSION;type:integer;nullable:NO;default:1;not null"  json:"recversion"`
	PARTITION          int64   `gorm:"column:PARTITION;default:5637144576;not null"   json:"partition"`
	RECID              int64   `gorm:"column:RECID;type:bigint;primaryKey;not null"  json:"recid"`
}

type EtaxTransRepository interface {
	GetById(string, string) ([]EtaxTransRepositoryDb, error)
}
