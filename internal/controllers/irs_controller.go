package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

type IRSController struct{}

func NewIRSController() *IRSController {
	return &IRSController{}
}

func (ctrl IRSController) Transaction(c *fiber.Ctx) error {
	var query dto.IrsTransactionReq
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.IrsTransactionRes{
			Success: false,
			Msg:     err.Error(),
		})
	}

	return c.JSON(dto.IrsTransactionRes{
		Success: true,
		Produk:  query.KodeProduk,
		Tujuan:  query.Tujuan,
		ReffID:  utils.RandString("1234567890", 10),
		RC:      "0068",
		Msg:     "Under proses",
	})
}
