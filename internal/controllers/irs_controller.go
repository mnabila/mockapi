package controllers

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

type IRSController struct {
	notifyUrl string
}

func NewIRSController(notify string) *IRSController {
	return &IRSController{notifyUrl: notify}
}

func (ctrl IRSController) Transaction(c *fiber.Ctx) error {
	var query dto.IrsTransactionReq
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.IrsTransactionRes{
			Success: false,
			Msg:     err.Error(),
		})
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		notify, err := url.Parse(ctrl.notifyUrl)
		if err != nil {
			return
		}

		value := url.Values{}
		value.Set("serverid", utils.RandString("1234567890", 15))
		value.Set("clientid", query.IDTrx)
		value.Set("statuscode", "1")
		value.Set("kp", query.KodeProduk)
		value.Set("msisdn", query.Tujuan)
		value.Set("sn", utils.RandString("1234567890", 12))
		value.Set("ms", fmt.Sprintf("TRX %s.%s BERHASIL", query.KodeProduk, query.Tujuan))
		notify.RawQuery = value.Encode()

		agent := fiber.Get(notify.String())
		agent.Bytes()
	}()

	return c.JSON(dto.IrsTransactionRes{
		Success: true,
		Produk:  query.KodeProduk,
		Tujuan:  query.Tujuan,
		ReffID:  utils.RandString("1234567890", 10),
		RC:      "0068",
		Msg:     "Under proses",
	})
}
