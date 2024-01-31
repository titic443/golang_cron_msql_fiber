package repository

import (
	"github.com/hirochachacha/go-smb2"
)

type fileshareRepository struct {
	client *smb2.Session
}

func NewfileshareRepository(client *smb2.Client) fileshareRepository {
	return fileshareRepository{client: client}
}

func (rp fileshareRepository) DownloadFile(remote string) error {
	return nil
}
