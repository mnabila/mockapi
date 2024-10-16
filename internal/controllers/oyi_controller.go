package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/dto"
	"github.com/mnabila/mockapi/internal/pkg/cache"
	"github.com/mnabila/mockapi/internal/utils"
)

type OYTController struct {
	cache     *cache.Cache
	notifyUrl string
}

func NewOYTController(notify string) *OYTController {
	return &OYTController{
		cache:     cache.NewCache(),
		notifyUrl: notify,
	}
}

func (ctrl OYTController) CreateEwallet(c *fiber.Ctx) error {
	var body dto.OYICreateEwalletReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.OYIEwalletRes{
			Status: dto.OYIStatus{
				Code:    "990",
				Message: "Request is Rejected (Parameter is invalid)",
			},
		})
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
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

	result := dto.OYIEwalletRes{
		Status: dto.OYIStatus{
			Code:    "000",
			Message: "success",
		},
		EwalletTrxStatus: "WAITING_PAYMENT",
		Amount:           body.Amount,
		TrxId:            utils.RandString("1234567890asdfghjklqwertyuiopzxcvbnm", 32),
		CustomerId:       body.CustomerId,
		PartnerTrxId:     body.PartnerTrxId,
		EwalletCode:      body.EwalletCode,
		EwalletURL:       body.SuccessRedirectURL,
	}

	ttl := 10 * time.Second
	ctrl.cache.Set(body.PartnerTrxId, result, ttl)

	return c.JSON(result)
}

func (ctrl OYTController) StatusEwallet(c *fiber.Ctx) error {
	var body dto.OYIStatusEwalletReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.OYIEwalletRes{
			Status: dto.OYIStatus{
				Code:    "990",
				Message: "Request is Rejected (Parameter is invalid)",
			},
		})
	}

	cache, ok := ctrl.cache.Get(body.PartnerTrxId)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.OYIEwalletRes{
			Status: dto.OYIStatus{
				Code:    "990",
				Message: "Request is Rejected (Parameter is invalid)",
			},
		})
	}

	trx, ok := cache.(dto.OYIEwalletRes)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.OYIEwalletRes{
			Status: dto.OYIStatus{
				Code:    "990",
				Message: "Request is Rejected (Parameter is invalid)",
			},
		})
	}

	return c.JSON(dto.OYIEwalletRes{
		Status: dto.OYIStatus{
			Code:    "000",
			Message: "success",
		},
		EwalletTrxStatus: "WAITING_PAYMENT",
		Amount:           trx.Amount,
		TrxId:            trx.TrxId,
		CustomerId:       trx.CustomerId,
		PartnerTrxId:     trx.PartnerTrxId,
		EwalletCode:      trx.EwalletCode,
		EwalletURL:       trx.EwalletURL,
	})
}
