package repository

import (
	"encoding/base64"
	"fmt"
	"go-etax/internal/logs"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/hirochachacha/go-smb2"
)

type fileshareRepository struct {
	client  *smb2.Session
	share   string
	workdir string
}

func NewfileshareRepository(client *smb2.Client, share string, workdir string) *fileshareRepository {
	return &fileshareRepository{client: client, share: share, workdir: workdir}
}

func (rp *fileshareRepository) ListFile() error {
	fmt.Println("List file")
	fs, err := rp.client.Mount(rp.share)
	if err != nil {
		logs.Error(err)
		return err
	}
	files, err := fs.ReadDir(rp.workdir)
	if err != nil {
		logs.Error(err)
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println("LIst file end")
	return nil
}

func (rp *fileshareRepository) DownloadFile(doc string) error {
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
				return err
			}
			v2 := fmt.Sprintf("./download/%v", file.Name())
			_ = v2
			lfs, err := os.Create(v2)
			if err != nil {
				return err
			}

			_, err = lfs.Write(fileByte)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (rp *fileshareRepository) DecodeFile(doc string) (*string, error) {
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
				return nil, err
			}

			rawEncode := base64.StdEncoding.EncodeToString(fileByte)
			o = rawEncode
			// return &rawDecode, nil

		}
	}
	log.Info("ENCODE PDF FILE SUCCESS " + doc)
	return &o, nil
}
