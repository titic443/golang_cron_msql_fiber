package repository

import "time"

type EtaxTable struct {
	ADJUST_AMOUNT_UNTAXED   float32   `gorm:"column:ADJUST_AMOUNT_UNTAXED;type:numeric(32,16);not null" json:"adjust_amount_untaxed"`
	BUYER_ADDRESS_LINE1     string    `gorm:"column:BUYER_ADDRESS_LINE1;type:nvarchar(250);not null" json:"buyer_address_line1"`
	BUYER_ADDRESS_LINE2     string    `gorm:"column:BUYER_ADDRESS_LINE2;type:nvarchar(250);not null" json:"buyer_address_line2"`
	BUYER_ADDRESS_LINE3     string    `gorm:"column:BUYER_ADDRESS_LINE3;type:nvarchar(250);not null" json:"buyer_address_line3"`
	BUYER_ADDRESS_LINE4     string    `gorm:"column:BUYER_ADDRESS_LINE4;type:nvarchar(250);not null" json:"buyer_address_line4"`
	BUYER_ADDRESS_LINE5     string    `gorm:"column:BUYER_ADDRESS_LINE5;type:nvarchar(250);not null" json:"buyer_address_line5"`
	BUYER_BRANCH_ID         string    `gorm:"column:BUYER_BRANCH_ID;type:nvarchar(5);not null" json:"buyer_branch_id"`
	BUYER_BUILDING_NAME     string    `gorm:"column:BUYER_BUILDING_NAME;type:nvarchar(100);not null" json:"buyer_building_name"`
	BUYER_BUILDING_NO       string    `gorm:"column:BUYER_BUILDING_NO;type:nvarchar(20);not null" json:"buyer_building_no"`
	BUYER_CITY_NAME         string    `gorm:"column:BUYER_CITY_NAME;type:nvarchar(100);not null" json:"buyer_city_name"`
	BUYER_COUNTRY_CODE      string    `gorm:"column:BUYER_COUNTRY_CODE;type:nvarchar(10);not null" json:"buyer_country_code"`
	BUYER_EMAIL             string    `gorm:"column:BUYER_EMAIL;type:nvarchar(100);not null" json:"buyer_email"`
	BUYER_NAME              string    `gorm:"column:BUYER_NAME;type:nvarchar(130);not null" json:"buyer_name"`
	BUYER_REF_DOCUMENT      string    `gorm:"column:BUYER_REF_DOCUMENT;type:nvarchar(50);not null" json:"buyer_ref_document"`
	BUYER_TAX_ID            string    `gorm:"column:BUYER_TAX_ID;type:nvarchar(15);not null" json:"buyer_tax_id"`
	BUYER_ZIP               string    `gorm:"column:BUYER_ZIP;type:nvarchar(10);not null" json:"buyer_zip"`
	COMPANY                 string    `gorm:"column:COMPANY;type:nvarchar(10);not null" json:"company"`
	CREATE_PURPOSE          string    `gorm:"column:CREATE_PURPOSE;type:nvarchar(10);not null" json:"create_purpose"`
	CREATE_PURPOSE_CODE     string    `gorm:"column:CREATE_PURPOSE_CODE;type:nvarchar(10);not null" json:"create_purpose_code"`
	CURRENCY_CODE           string    `gorm:"column:CURRENCY_CODE;type:nvarchar(10);not null" json:"currency_code"`
	DOCUMENT_ID             string    `gorm:"column:DOCUMENT_ID;type:nvarchar(20);not null" json:"document_id"`
	DOCUMENT_ISSUE_DTM      time.Time `gorm:"column:DOCUMENT_ISSUE_DTM;type:datetime;not null" json:"document_issue_dtm"`
	DOCUMENT_TYPE_CODE      string    `gorm:"column:DOCUMENT_TYPE_CODE;type:nvarchar(10);not null" json:"document_type_code"`
	FINAL_AMOUNT_UNTAXED    float32   `gorm:"column:FINAL_AMOUNT_UNTAXED;type:numeric(32,16);not null" json:"final_amount_untaxed"`
	FORM_NAME               string    `gorm:"column:FORM_NAME;type:nvarchar(10);not null" json:"form_name"`
	FORM_TYPE               string    `gorm:"column:FORM_TYPE;type:nvarchar(20);not null" json:"form_type"`
	ORIGINAL_AMOUNT_UNTAXED float32   `gorm:"column:ORIGINAL_AMOUNT_UNTAXED;type:numeric(32,16);not null" json:"original_amount_untaxed"`
	PDF_CONTENT             string    `gorm:"column:PDF_CONTENT;type:nvarchar(10);not null" json:"pdf_content"`
	REF_DOCUMENT_ID         string    `gorm:"column:REF_DOCUMENT_ID;type:nvarchar(20);not null" json:"ref_document_id"`
	REF_DOCUMENT_ISSUE_DTM  string    `gorm:"column:REF_DOCUMENT_ISSUE_DTM;type:datetime;not null" json:"ref_document_issue_dtm"`
	REF_DOCUMENT_TYPE_CODE  string    `gorm:"column:REF_DOCUMENT_TYPE_CODE;type:nvarchar(10);not null" json:"ref_document_type_code"`
	SELLER_BRANCH_ID        string    `gorm:"column:SELLER_BRANCH_ID;type:nvarchar(5);not null" json:"seller_branch_id"`
	SELLER_TAX_ID           string    `gorm:"column:SELLER_TAX_ID;type:nvarchar(15);not null" json:"seller_tax_id"`
	SEND_MAIL               string    `gorm:"column:SEND_MAIL;type:nvarchar(10);not null" json:"send_mail"`
	SOURCE_SYSTEM           string    `gorm:"column:SOURCE_SYSTEM;type:nvarchar(300);not null" json:"source_system"`
	STATUS_SIGN             int       `gorm:"column:STATUS_SIGN;type:integer;not null" json:"status_sign"`
	RECVERSION              int       `gorm:"column:RECVERSION;type:integer;nullable:NO;default:1;not null" json:"recversion"`
	PARTITION               int64     `gorm:"column:PARTITION;default:5637144576;not null" json:"partition"`
	RECID                   int64     `gorm:"column:RECID;type:bigint;primaryKey;not null" json:"recid"`
}

type EtaxTableRepository interface {
	SqlGetAll() ([]EtaxTable, error)
	SqlGetById(int) (*EtaxTable, error)
	SqlUpdate(int) error
}
