package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mnabila/mockapi/internal/configs"
	"github.com/mnabila/mockapi/internal/routers"
)

func main() {
	conf, err := configs.NewEnvConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		ServerHeader:            "MockAPI.Service",
		AppName:                 "MockAPI.Service",
		EnableTrustedProxyCheck: true,
	})

	app.Use(recover.New())

	routers.UseFinpay(app)
	routers.UseIrs(app)
	routers.UseOYI(app)

	if err := app.Listen(conf.GetListenAddress()); err != nil {
		log.Fatalf("failed run service in %s with error message: %v", conf.GetListenAddress(), err)
	}
}
