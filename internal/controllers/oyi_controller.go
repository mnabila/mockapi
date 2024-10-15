package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/utils"
)

type OYTController struct {
	notifyUrl string
}

func NewOYTController(notify string) *OYTController {
	return &OYTController{
		notifyUrl: notify,
	}
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

	go func() {
		timestamp := time.Now().Add(24 * time.Hour).Format("02/01/2006T15:04:05.000-0700")
		payload := dto.OYIEwalletNotify{
			Success:            true,
			PartnerTrxId:       body.PartnerTrxId,
			TrxId:              utils.RandString("1234567890qwertyuiop", 32),
			CustomerId:         body.CustomerId,
			Amount:             body.Amount,
			EwalletCode:        body.EwalletCode,
			MobileNumber:       body.MobileNumber,
			SuccessRedirectURL: body.SuccessRedirectURL,
			SettlementTime:     timestamp,
			SettlementStatus:   "WAITING",
		}

		agent := fiber.Post(ctrl.notifyUrl)
		agent.JSON(payload)
		agent.Bytes()
	}()

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
