package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/configs"
	"github.com/mnabila/mockapi/internal/controllers"
)

func UseIrs(conf *configs.EnvConfig, r fiber.Router) {
	ctrl := controllers.NewIRSController(conf.IRSNotifyUrl)
	r.Get("/api/h2h", ctrl.Transaction)
}
