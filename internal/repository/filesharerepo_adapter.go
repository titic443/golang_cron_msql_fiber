package repository

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hirochachacha/go-smb2"
)

type fileshareRepository struct {
	client  *smb2.Session
	share   string
	workdir string
}

func NewfileshareRepository(client *smb2.Client, share string, workdir string) fileshareRepository {
	return fileshareRepository{client: client, share: share, workdir: workdir}
}

func (rp fileshareRepository) DownloadFile(doc string) error {
	fs, err := rp.client.Mount(rp.share)
	if err != nil {
		return err
	}

	files, err := fs.ReadDir(rp.workdir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), doc) {
			v1 := fmt.Sprintf("%v/%v", rp.workdir, file.Name())
			fileByte, _ := fs.ReadFile(v1)
			if err != nil {
				log.Println(err)
			}
			v2 := fmt.Sprintf("./download/%v", file.Name())
			_ = v2
			lfs, err := os.Create(v2)
			if err != nil {
				log.Println(err)
			}

			_, err = lfs.Write(fileByte)
			if err != nil {
				log.Println(err)
			}

		}
	}

	return nil
}

func (rp fileshareRepository) DecodeFile(doc string) (*string, error) {
	var o string
	fs, err := rp.client.Mount(rp.share)
	if err != nil {
		return nil, err
	}

	files, err := fs.ReadDir(rp.workdir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), doc) {
			v1 := fmt.Sprintf("%v/%v", rp.workdir, file.Name())
			fileByte, _ := fs.ReadFile(v1)
			if err != nil {
				log.Println(err)
			}

			rawEncode := base64.StdEncoding.EncodeToString(fileByte)
			o = rawEncode
			// return &rawDecode, nil

		}
	}

	return &o, nil
}
