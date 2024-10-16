package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/configs"
	"github.com/mnabila/mockapi/internal/controllers"
)

func UseOYI(conf *configs.EnvConfig, r fiber.Router) {
	ctrl := controllers.NewOYTController(conf.OYINotifyUrl)
	r.Post("/api/e-wallet-aggregator/create-transaction", ctrl.Ewallet)
}
