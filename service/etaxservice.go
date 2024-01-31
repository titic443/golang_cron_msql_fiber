package service

import (
	"time"
)

type ResponseData struct {
	FORM_TYPE   string  `json:"form_type"`
	FORM_NAME   string  `json:"form_name"`
	PDF_CONTENT string  `json:"pdf_content"`
	DocData     DocData `json:"doc_data"`
}

type DocData struct {
	SOURCE_SYSTEM           string                `json:"source_system"`
	SEND_MAIL               string                `json:"send_mail"`
	CURRENCY_CODE           string                `json:"currency_code"`
	DOCUMENT_TYPE_CODE      string                `json:"document_type_code"`
	DOCUMENT_ID             string                `json:"document_id"`
	DOCUMENT_ISSUE_DTM      time.Time             `json:"document_issue_dtm"`
	CREATE_PURPOSE_CODE     string                `json:"create_purpose_code"`
	CREATE_PURPOSE          string                `json:"create_purpose"`
	REF_DOCUMENT_ID         string                `json:"ref_document_id"`
	REF_DOCUMENT_ISSUE_DTM  string                `json:"ref_document_issue_dtm"`
	REF_DOCUMENT_TYPE_CODE  string                `json:"ref_document_type_code"`
	BUYER_REF_DOCUMENT      string                `json:"buyer_ref_document"`
	SELLER_BRANCH_ID        string                `json:"seller_branch_id"`
	SELLER_TAX_ID           string                `json:"seller_tax_id"`
	BUYER_NAME              string                `json:"buyer_name"`
	BuyerType               string                `json:"buyer_type"`
	BUYER_TAX_ID            string                `json:"buyer_tax_id"`
	BUYER_BRANCH_ID         string                `json:"buyer_branch_id"`
	BUYER_EMAIL             string                `json:"buyer_email"`
	BUYER_ZIP               string                `json:"buyer_zip"`
	BUYER_BUILDING_NAME     string                `json:"buyer_building_name"`
	BUYER_BUILDING_NO       string                `json:"buyer_building_no"`
	BUYER_ADDRESS_LINE1     string                `json:"buyer_address_line1"`
	BUYER_ADDRESS_LINE2     string                `json:"buyer_address_line2"`
	BUYER_ADDRESS_LINE3     string                `json:"buyer_address_line3"`
	BUYER_ADDRESS_LINE4     string                `json:"buyer_address_line4"`
	BUYER_ADDRESS_LINE5     string                `json:"buyer_address_line5"`
	BUYER_CITY_NAME         string                `json:"buyer_city_name"`
	BUYER_COUNTRY_CODE      string                `json:"buyer_country_code"`
	LineItemInformation     []LineItemInformation `json:"line_item_information"`
	ORIGINAL_AMOUNT_UNTAXED int                   `json:"original_amount_untaxed"`
	FINAL_AMOUNT_UNTAXED    int                   `json:"final_amount_untaxed"`
	ADJUST_AMOUNT_UNTAXED   int                   `json:"adjust_amount_untaxed"`
}

type LineItemInformation struct {
	PRODUCT_CODE       string  `json:"product_code"`
	PRODUCT_NAME       string  `json:"product_name"`
	PRODUCT_PRICE      float64 `json:"line_tax_type_code"`
	PRODUCT_QUANTITY   int     `json:"product_price"`
	LINE_TAX_TYPE_CODE string  `json:"product_quantity"`
	LINE_TAX_RATE      int     `json:"line_tax_rate"`
	LINE_BASE_AMOUNT   int     `json:"line_base_amount"`
	LINE_TAX_AMOUNT    int     `json:"line_tax_amount"`
	LINE_TOTAL_AMOUNT  int     `json:"line_total_amount"`
}

type EtaxService interface {
	SignEtax() ([]ResponseData, error)
}
