package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-etax/internal/logs"
	"go-etax/internal/service"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"

	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

type etaxTableHandler struct {
	etaxTableSrv service.EtaxService
}

func NewEtaxTableHandler(etaxTableSrv service.EtaxService) etaxTableHandler {

	return etaxTableHandler{etaxTableSrv: etaxTableSrv}
}

func (h etaxTableHandler) SendEtaxToEco(c *fiber.Ctx) error {
	url := viper.GetString("api.url")
	token := fmt.Sprintf("token %s", viper.GetString("api.token"))
	o, err := h.etaxTableSrv.SignEtax()
	if err != nil {
		logs.Error(err)
	}
	for _, v := range o {
		obyte, _ := json.Marshal(v)

		r, _ := http.NewRequest(fiber.MethodPost, url, bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", token)
		client := &http.Client{}
		http.DefaultClient.Timeout = 5 * time.Second
		res, err := client.Do(r)
		if err != nil {
			logs.Error(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			b, _ := io.ReadAll(res.Body)
			logs.Error(b)

		} else {
			h.etaxTableSrv.SqlUpdateSuccess(v.DocData.DOCUMENT_ID)
		}
	}

	return c.JSON(o)
}

func (h etaxTableHandler) SendEtaxToEcoCronjob(cronEntries ...[]cron.Entry) error {
	logs.Info("Start call DB and Ecosoft")
	url := viper.GetString("api.url")
	token := fmt.Sprintf("token %s", viper.GetString("api.token"))
	o, err := h.etaxTableSrv.SignEtax()
	if err != nil {
		log.Error(err)
	}
	for _, v := range o {
		obyte, _ := json.Marshal(v)

		r, _ := http.NewRequest(fiber.MethodPost, url, bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", token)
		client := &http.Client{}
		http.DefaultClient.Timeout = 5 * time.Second
		res, err := client.Do(r)

		if err != nil {
			logs.Error(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			b, _ := io.ReadAll(res.Body)
			logs.Error(b)

		} else {
			h.etaxTableSrv.SqlUpdateSuccess(v.DocData.DOCUMENT_ID)
		}
	}
	return nil
}
