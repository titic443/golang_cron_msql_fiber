package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-etax/internal/service"

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

		fmt.Println(res)
	}

	return c.JSON(o)
}
