package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/controllers"
)

func UseFinpay(r fiber.Router) {
	ctrl := controllers.NewFinpayController()
	r.Post("/pg/payment/card/initiate", ctrl.Ewallet)
}

