package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-etax/internal/logs"
	"go-etax/internal/repository"
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
	var docId string
	url := viper.GetString("api.url")
	token := fmt.Sprintf("token %s", viper.GetString("api.token"))
	o, err := h.etaxTableSrv.SignEtax()
	if err != nil {
		logs.Error(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	for _, v := range o {
		mapOutput := make(map[string]interface{})
		obyte, _ := json.Marshal(v)

		err = json.Unmarshal(obyte, &mapOutput)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
		}
		for keyOutput, valueOutput := range mapOutput {
			if keyOutput == "doc_data" {
				docData := repository.EtaxTable{}
				jsonDocdata, _ := json.Marshal(valueOutput)
				err = json.Unmarshal(jsonDocdata, &docData)
				if err != nil {
					fmt.Println("Error unmarshaling JSON:", err)
				}
				docId = docData.DOCUMENT_ID
			}

		}

		r, _ := http.NewRequest(fiber.MethodPost, url, bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", token)
		// client := &http.Client{}
		client := &http.Client{Transport: tr}
		http.DefaultClient.Timeout = 5 * time.Second
		res, err := client.Do(r)
		if err != nil {
			logs.Error(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			var trace map[string]interface{}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)

			}
			if err := json.Unmarshal(b, &trace); err != nil {
				fmt.Println("Error unmarshaling JSON:", err)

			}

			var resMsg = make(map[string]interface{})
			v, ok := trace["_server_messages"]
			if ok {

				resMsg["err"] = v
				resMsg["document_id"] = docId
				logs.Error(resMsg)

			} else {

				logs.Error(string(b))
			}

		} else {
			h.etaxTableSrv.SqlUpdateSuccess(v.DocData.DOCUMENT_ID)
		}
	}

	return c.JSON(o)
}

func (h etaxTableHandler) SendEtaxToEcoCronjob(cronEntries ...[]cron.Entry) error {
	var docId string
	logs.Info("Start call DB and Ecosoft")
	url := viper.GetString("api.url")
	token := fmt.Sprintf("token %s", viper.GetString("api.token"))
	o, err := h.etaxTableSrv.SignEtax()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if err != nil {
		log.Error(err)
	}
	for _, v := range o {
		mapOutput := make(map[string]interface{})
		obyte, _ := json.Marshal(v)
		err = json.Unmarshal(obyte, &mapOutput)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
		}
		for keyOutput, valueOutput := range mapOutput {
			if keyOutput == "doc_data" {
				docData := repository.EtaxTable{}
				jsonDocdata, _ := json.Marshal(valueOutput)
				err = json.Unmarshal(jsonDocdata, &docData)
				if err != nil {
					fmt.Println("Error unmarshaling JSON:", err)
				}
				docId = docData.DOCUMENT_ID
			}

		}
		r, _ := http.NewRequest(fiber.MethodPost, url, bytes.NewBuffer(obyte))

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", token)
		client := &http.Client{Transport: tr}
		http.DefaultClient.Timeout = 5 * time.Second
		res, err := client.Do(r)

		if err != nil {
			logs.Error(err)
		}
		// defer res.Body.Close()
		if res.StatusCode != 200 {
			var trace map[string]interface{}
			b, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)

			}
			if err := json.Unmarshal(b, &trace); err != nil {
				fmt.Println("Error unmarshaling JSON:", err)

			}
			var resMsg = make(map[string]interface{})
			v, ok := trace["_server_messages"]
			if ok {

				resMsg["err"] = v
				resMsg["document_id"] = docId
				logs.Error(resMsg)

			} else {

				logs.Error(string(b))
			}
		} else {
			h.etaxTableSrv.SqlUpdateSuccess(v.DocData.DOCUMENT_ID)
		}
	}
	return nil
}
