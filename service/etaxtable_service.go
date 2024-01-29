package service

import "go-etax/repository"

type etaxTableService struct {
	etaxRepo repository.EtaxTableRepository
}

func NewEtaxTableService(etaxRepo repository.EtaxTableRepository) etaxTableService {
	return etaxTableService{etaxRepo: etaxRepo}
}
