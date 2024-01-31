package handler

import (
	"go-etax/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type etaxTableHandler struct {
	etaxTableSrv service.EtaxTableService
}

type Test struct {
	STATUS_SIGN int   `json:"status_sign"`
	RECVERSION  int   `json:"recversion"`
	PARTITION   int64 `json:"partition"`
	RECID       int64 `json:"recid"`
}

func NewEtaxTableHandler(etaxTableSrv service.EtaxTableService) etaxTableHandler {
	return etaxTableHandler{etaxTableSrv: etaxTableSrv}
}

func (h etaxTableHandler) SendEtaxToEco(c *fiber.Ctx) error {
	r, err := h.etaxTableSrv.GetEtaxTable()
	if err != nil {
		log.Println(err)
	}

	return c.JSON(r)
}
