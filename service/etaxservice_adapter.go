package service

import (
	"go-etax/repository"
	"log"

	"github.com/mitchellh/mapstructure"
)

type etaxService struct {
	etaxTableRepo repository.EtaxTableRepository
	etaxTransRepo repository.EtaxTransRepository
}

func NewEtaxTableService(etaxTableRepo repository.EtaxTableRepository, etaxTransRepo repository.EtaxTransRepository) etaxService {
	return etaxService{etaxTableRepo: etaxTableRepo, etaxTransRepo: etaxTransRepo}
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
		responses = append(responses, response)
	}
	return responses, nil
}
