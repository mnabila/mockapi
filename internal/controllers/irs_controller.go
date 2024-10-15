package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

func IrsController(c *fiber.Ctx) error {
	var query dto.IrsReq
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.IrsRes{
			Success: false,
			Msg:     err.Error(),
		})
	}

	return c.JSON(dto.IrsRes{
		Success: true,
		Produk:  query.KodeProduk,
		Tujuan:  query.Tujuan,
		ReffID:  utils.RandString("1234567890", 10),
		RC:      "0068",
		Msg:     "Under proses",
	})
}
