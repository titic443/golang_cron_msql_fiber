package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-etax/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type etaxTableHandler struct {
	etaxTableSrv service.EtaxService
}

func NewEtaxTableHandler(etaxTableSrv service.EtaxService) etaxTableHandler {
	return etaxTableHandler{etaxTableSrv: etaxTableSrv}
}

func (h etaxTableHandler) SendEtaxToEco(c *fiber.Ctx) error {
	o, err := h.etaxTableSrv.SignEtax()
	if err != nil {
		log.Error(err)
	}

	obyte, _ := json.Marshal(o)
	r, _ := http.NewRequest(fiber.MethodPost, "http://localhost:3000/etax", bytes.NewBuffer(obyte))

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer fmt.Println("End", res)
	return nil
}
