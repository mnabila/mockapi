package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

func OYIController(c *fiber.Ctx) error {
	var body dto.OYITransactionReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.IrsRes{
			Success: false,
			Msg:     err.Error(),
		})
	}

	return c.JSON(dto.OYITransactionRes{
		Status: dto.OYIStatus{
			Code:    "000",
			Message: "success",
		},
		EwalletTrxStatus: "WAITING_PAYMENT",
		Amount:           body.Amount,
		TrxID:            utils.RandString("1234567890asdfghjklqwertyuiopzxcvbnm", 32),
		CustomerID:       body.CustomerId,
		PartnerTrxID:     body.PartnerTrxId,
		EwalletCode:      body.EwalletCode,
		EwalletURL:       body.SuccessRedirectURL,
	})
}
