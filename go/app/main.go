package main

import (
	"log"

	"./routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New(), logger.New())
	key := "tokenKey"
	routers.Version1(app, key)
	log.Fatal(app.Listen(":3000"))
}
