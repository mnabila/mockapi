package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/configs"
	"github.com/mnabila/mockapi/internal/controllers"
)

func UseOYI(conf *configs.EnvConfig, r fiber.Router) {
	ctrl := controllers.NewOYTController(conf.FinpayNotifyUrl)
	r.Post("/api/h2h", ctrl.Ewallet)
}
