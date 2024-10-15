package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
)

type FinpayController struct{}

func NewFinpayController() *FinpayController {
	return &FinpayController{}
}

func (ctrl FinpayController) Ewallet(c *fiber.Ctx) error {
	start := time.Now()
	var body dto.FinpayEwalletReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.FinpayEwalletRes{
			ResponseCode:    "4000000",
			ResponseMessage: err.Error(),
		})
	}

	return c.JSON(dto.FinpayEwalletRes{
		ResponseCode:    "2000000",
		ResponseMessage: "Successfull",
		RedirectURL:     body.URL.CallbackURL,
		ExpiryLink:      time.Now().Add(time.Duration(body.Order.Timeout) * time.Minute).Format("2006-01-02 15:04:05"),
		PaymentCode:     nil,
		ProcessingTime:  time.Since(start).Seconds(),
	})
}
