package controllers

import (
	"fmt"
	"net/url"

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
		notify, err := url.Parse(ctrl.notifyUrl)
		if err != nil {
			return
		}

		value := url.Values{}
		value.Set("id", query.ID)
		value.Set("pin", query.PIN)
		value.Set("user", query.User)
		value.Set("pass", query.Pass)
		value.Set("kodeproduk", query.KodeProduk)
		value.Set("tujuan", query.KodeProduk)
		value.Set("counter", fmt.Sprint(query.Counter))
		value.Set("idtrx", query.IDTrx)
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
