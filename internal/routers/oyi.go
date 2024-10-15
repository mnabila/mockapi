package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mockapi/internal/controllers"
)

func UseOYI(r fiber.Router) {
	ctrl := controllers.NewIRSController()
	r.Post("/api/h2h", ctrl.Transaction)
}
