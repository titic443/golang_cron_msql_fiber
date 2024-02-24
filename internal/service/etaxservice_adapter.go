package service

import (
	"fmt"
	"go-etax/internal/logs"
	"go-etax/internal/repository"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type etaxService struct {
	etaxTableRepo repository.EtaxTableRepository
	etaxTransRepo repository.EtaxTransRepository
	fileshareRepo repository.FileshareRepository
}

func NewEtaxTableService(etaxTableRepo repository.EtaxTableRepository, etaxTransRepo repository.EtaxTransRepository, fileshareRepo repository.FileshareRepository) *etaxService {
	return &etaxService{etaxTableRepo: etaxTableRepo, etaxTransRepo: etaxTransRepo, fileshareRepo: fileshareRepo}
}

func (s *etaxService) SignEtax() ([]ResponseData, error) {
	etaxTables, err := s.etaxTableRepo.SqlGetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if len(etaxTables) != 0 {
		var responses []ResponseData
		var lineInformations []LineItemInformation
		for _, etaxTable := range etaxTables {
			var response ResponseData
			var docData DocData
			if mapstructure.Decode(etaxTable, &response); err != nil {
				logs.Error(err)
				// return nil, err
			}
			fmt.Println(etaxTable.DOCUMENT_ID)
			o, err := s.EncodePdf(etaxTable.DOCUMENT_ID)
			if err != nil {
				logs.Error(err)
				// return nil, err
			}
			if o != nil {

				response.PDF_CONTENT = *o
			} else {
				response.PDF_CONTENT = ""
			}

			rs, err := s.etaxTransRepo.GetById(etaxTable.DOCUMENT_ID, etaxTable.COMPANY)
			if err != nil {
				logs.Error(err)
				continue
			}
			for _, r := range rs {
				var lineInformation LineItemInformation
				if mapstructure.Decode(r, &lineInformation); err != nil {
					logs.Error(err)
					continue
				}
				lineInformations = append(lineInformations, lineInformation)
			}

			docData.LineItemInformation = lineInformations
			if mapstructure.Decode(etaxTable, &docData); err != nil {
				logs.Error(err)
				continue
			}
			response.DocData = docData
			t, _ := s.Transform(&response.DocData)
			_ = t

			// if err = s.etaxTableRepo.SqlUpdate(&etaxTable); err != nil {
			// 	return nil, err
			// }

			responses = append(responses, response)
		}

		return responses, nil
	}
	logs.Debug("No avilable document to sign ETAX")
	return nil, nil
}

func (s *etaxService) Transform(p *DocData) (*DocData, error) {
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
	return p, nil
}

func (s *etaxService) EncodePdf(docId string) (*string, error) {
	o, err := s.fileshareRepo.DecodeFile(docId)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s *etaxService) SqlUpdateSuccess(docId string) error {
	if err := s.etaxTableRepo.SqlUpdate(docId); err != nil {
		return err
	}
	return nil
}

func (s *etaxService) ListFile() {
	s.fileshareRepo.ListFile()
}
