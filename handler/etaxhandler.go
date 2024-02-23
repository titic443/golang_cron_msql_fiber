package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-etax/internal/logs"
	"go-etax/internal/service"
	"io"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
		log.Error(err)
	}
	for _, v := range o {
		obyte, _ := json.Marshal(v)

		r, _ := http.NewRequest(fiber.MethodPost, url, bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", token)
		client := &http.Client{}
		res, err := client.Do(r)
		res.Body.Close()
		if err != nil {
			panic(err)
		}

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
	o, err := h.etaxTableSrv.SignEtax()
	if err != nil {
		log.Error(err)
	}
	for _, v := range o {
		obyte, _ := json.Marshal(v)

		r, _ := http.NewRequest(fiber.MethodPost, "https://etax-uat.energyabsolute.co.th/api/method/etax_inet.api.etax.sign_etax_document", bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", "token c748b18630bd8f6:954599aeaa941d7")
		client := &http.Client{}
		res, err := client.Do(r)

		if err != nil {
			panic(err)
		}

		if res.StatusCode != 200 {
			b, _ := io.ReadAll(res.Body)
			logs.Error(b)

		} else {
			h.etaxTableSrv.SqlUpdateSuccess(v.DocData.DOCUMENT_ID)
		}
	}
	return nil
}
