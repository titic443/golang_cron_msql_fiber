package service

import (
	"encoding/base64"
	"fmt"
	"go-etax/internal/logs"
	"net"
	"strings"

	"github.com/hirochachacha/go-smb2"
)

type User struct {
	serverIP     string
	userName     string
	userPassword string
	shareName    string
	folder       string
}

func connectSMBserver(user User) (*smb2.Session, net.Conn) {
	conn, err := net.Dial("tcp", user.serverIP+":445")
	if err != nil {
		panic(err)
	}
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     user.userName,
			Password: user.userPassword,
		},
	}
	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	return s, conn
}

func getMount(s *smb2.Session, shareName string) *smb2.Share {
	// connect to share
	m, err := s.Mount(shareName)
	if err != nil {
		panic(err)
	}
	return m
}
func DecodeFile(s *smb2.Share, folder string, doc string) (*string, error) {
	var o string
	files, err := s.ReadDir(folder)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), doc) {
			v1 := fmt.Sprintf("%v/%v", folder, file.Name())
			fileByte, _ := s.ReadFile(v1)
			if err != nil {
				return nil, err
			}

			rawEncode := base64.StdEncoding.EncodeToString(fileByte)
			o = rawEncode

			logs.Info("ENCODE PDF FILE " + doc)
		}
	}
	return &o, nil
}
