package service

import (
	"go-etax/internal/repository"
	"log"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type etaxService struct {
	etaxTableRepo repository.EtaxTableRepository
	etaxTransRepo repository.EtaxTransRepository
	fileshareRepo repository.FileshareRepository
}

func NewEtaxTableService(etaxTableRepo repository.EtaxTableRepository, etaxTransRepo repository.EtaxTransRepository, fileshareRepo repository.FileshareRepository) etaxService {
	return etaxService{etaxTableRepo: etaxTableRepo, etaxTransRepo: etaxTransRepo, fileshareRepo: fileshareRepo}
}

func (s etaxService) SignEtax() ([]ResponseData, error) {
	etaxTables, err := s.etaxTableRepo.SqlGetAll()
	if err != nil {
		return nil, err
	}
	var responses []ResponseData
	var lineInformations []LineItemInformation
	for _, etaxTable := range etaxTables {
		var response ResponseData
		var docData DocData
		if mapstructure.Decode(etaxTable, &response); err != nil {
			log.Println(err)
		}
		o, err := s.EncodePdf(etaxTable.DOCUMENT_ID)
		if err != nil {
			log.Println(err)
		}
		response.PDF_CONTENT = *o

		rs, err := s.etaxTransRepo.GetById(etaxTable.DOCUMENT_ID, etaxTable.COMPANY)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, r := range rs {
			var lineInformation LineItemInformation
			if mapstructure.Decode(r, &lineInformation); err != nil {
				log.Println(err)
				continue
			}
			lineInformations = append(lineInformations, lineInformation)
		}

		docData.LineItemInformation = lineInformations
		if mapstructure.Decode(etaxTable, &docData); err != nil {
			log.Println(err)
			continue
		}
		response.DocData = docData
		s.Transform(&response.DocData)
		// b, f := strings.CutSuffix(response.DocData.DOCUMENT_ISSUE_DTM, "Z")
		// if f == true {

		// 	response.DocData.DOCUMENT_ISSUE_DTM = b
		// }
		// b, f = strings.CutSuffix(response.DocData.REF_DOCUMENT_ISSUE_DTM, "Z")
		// if f == true {
		// 	response.DocData.REF_DOCUMENT_ISSUE_DTM = b
		// }
		// if len(response.DocData.CREATE_PURPOSE) == 0 || len(response.DocData.CREATE_PURPOSE_CODE) == 0 {
		// 	response.DocData.REF_DOCUMENT_ID = ""
		// 	response.DocData.REF_DOCUMENT_ISSUE_DTM = ""
		// 	response.DocData.REF_DOCUMENT_TYPE_CODE = ""
		// }
		if err = s.etaxTableRepo.SqlUpdate(&etaxTable); err != nil {
			return nil, err
		}

		responses = append(responses, response)
	}
	return responses, nil
}

func (s etaxService) Transform(p *DocData) {
	if len(p.LineItemInformation) > 0 {
		for _, lineItem := range p.LineItemInformation {
			ok := strings.Contains(lineItem.LINE_TAX_TYPE_CODE, "VUD")
			if ok {
				lineItem.LINE_TAX_TYPE_CODE = "VAT"
			}
		}
	}
	b, ok := strings.CutSuffix(p.DOCUMENT_ISSUE_DTM, "Z")
	if ok {
		p.DOCUMENT_ISSUE_DTM = b
	}
	b, ok = strings.CutSuffix(p.REF_DOCUMENT_ISSUE_DTM, "Z")
	if ok {
		p.REF_DOCUMENT_ISSUE_DTM = b
	}
	if len(p.CREATE_PURPOSE) == 0 || len(p.CREATE_PURPOSE_CODE) == 0 {
		p.REF_DOCUMENT_ID = ""
		p.REF_DOCUMENT_ISSUE_DTM = ""
		p.REF_DOCUMENT_TYPE_CODE = ""
	}
}

func (s etaxService) EncodePdf(docId string) (*string, error) {
	o, err := s.fileshareRepo.DecodeFile(docId)
	if err != nil {
		return nil, err
	}
	return o, nil
}
