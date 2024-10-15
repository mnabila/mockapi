package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

type OYTController struct{}

func NewOYTController() *OYTController {
	return &OYTController{}
}

func (ctrl OYTController) Ewallet(c *fiber.Ctx) error {
	var body dto.OYIEwalletReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.OYIEwalletRes{
			Status: dto.OYIStatus{
				Code:    "990",
				Message: "	Request is Rejected (Parameter is invalid)",
			},
		})
	}

	return c.JSON(dto.OYIEwalletRes{
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
