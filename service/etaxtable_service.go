package service

import (
	"go-etax/repository"
	"log"

	"github.com/mitchellh/mapstructure"
)

type etaxTableService struct {
	etaxRepo repository.EtaxTableRepository
}

func NewEtaxTableService(etaxRepo repository.EtaxTableRepository) etaxTableService {
	return etaxTableService{etaxRepo: etaxRepo}
}

func (s etaxTableService) GetEtaxTable() ([]EtaxTableResponse, error) {
	etaxTables, err := s.etaxRepo.SqlGetAll()
	if err != nil {
		log.Println(err)
	}
	var responses []EtaxTableResponse
	// var jsonResArr []byte
	for _, etaxTable := range etaxTables {
		var response EtaxTableResponse
		if mapstructure.Decode(etaxTable, &response); err != nil {
			log.Println(err)
			continue
		}

		responses = append(responses, response)
		// jsonResArr, _ = json.Marshal(responses)

	}
	return responses, nil
}

func (s etaxTableService) UpdateEtaxTable(p *repository.EtaxTable) error {
	err := s.etaxRepo.SqlUpdate(p)
	if err != nil {
		return err
	}
	return nil
}
