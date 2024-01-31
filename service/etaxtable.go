package service

import "go-etax/repository"

type EtaxTableResponse struct {
	// ADJUST_AMOUNT_UNTAXED   float32   `json:"adjust_amount_untaxed"`
	// BUYER_ADDRESS_LINE1     string    `json:"buyer_address_line1"`
	// BUYER_ADDRESS_LINE2     string    `json:"buyer_address_line2"`
	// BUYER_ADDRESS_LINE3     string    `json:"buyer_address_line3"`
	// BUYER_ADDRESS_LINE4     string    `json:"buyer_address_line4"`
	// BUYER_ADDRESS_LINE5     string    `json:"buyer_address_line5"`
	// BUYER_BRANCH_ID         string    `json:"buyer_branch_id"`
	// BUYER_BUILDING_NAME     string    `json:"buyer_building_name"`
	// BUYER_BUILDING_NO       string    `json:"buyer_building_no"`
	// BUYER_CITY_NAME         string    `json:"buyer_city_name"`
	// BUYER_COUNTRY_CODE      string    `json:"buyer_country_code"`
	// BUYER_EMAIL             string    `json:"buyer_email"`
	// BUYER_NAME              string    `json:"buyer_name"`
	// BUYER_REF_DOCUMENT      string    `json:"buyer_ref_document"`
	// BUYER_TAX_ID            string    `json:"buyer_tax_id"`
	// BUYER_ZIP               string    `json:"buyer_zip"`
	// COMPANY                 string    `json:"company"`
	// CREATE_PURPOSE          string    `json:"create_purpose"`
	// CREATE_PURPOSE_CODE     string    `json:"create_purpose_code"`
	// CURRENCY_CODE           string    `json:"currency_code"`
	// DOCUMENT_ID             string    `json:"document_id"`
	// DOCUMENT_ISSUE_DTM      time.Time `json:"document_issue_dtm"`
	// DOCUMENT_TYPE_CODE      string    `json:"document_type_code"`
	// FINAL_AMOUNT_UNTAXED    float32   `json:"final_amount_untaxed"`
	// FORM_NAME               string    `json:"form_name"`
	// FORM_TYPE               string    `json:"form_type"`
	// ORIGINAL_AMOUNT_UNTAXED float32   `json:"original_amount_untaxed"`
	// PDF_CONTENT             string    `json:"pdf_content"`
	// REF_DOCUMENT_ID         string    `json:"ref_document_id"`
	// REF_DOCUMENT_ISSUE_DTM  string    `json:"ref_document_issue_dtm"`
	// REF_DOCUMENT_TYPE_CODE  string    `json:"ref_document_type_code"`
	// SELLER_BRANCH_ID        string    `json:"seller_branch_id"`
	// SELLER_TAX_ID           string    `json:"seller_tax_id"`
	// SEND_MAIL               string    `json:"send_mail"`
	// SOURCE_SYSTEM           string    `json:"source_system"`
	STATUS_SIGN int   `json:"status_sign"`
	RECVERSION  int   `json:"recversion"`
	PARTITION   int64 `json:"partition"`
	RECID       int64 `json:"recid"`
}

type EtaxTableService interface {
	GetEtaxTable() ([]EtaxTableResponse, error)
	UpdateEtaxTable(*repository.EtaxTable) error
}
