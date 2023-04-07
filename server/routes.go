package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Serve() {
	app := fiber.New()

	app.Static("/static", "./static")
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Minute * 5,
	})

	app.Get("/", Index)
	app.Post("/mobile-device", Query)

	app.Use(AuthMiddeware)
	app.Get("/form", Report)
	app.Post("/form", ReportIMEI)
	app.Get("/google_login", GoogleLogin)
	app.Get("/google_callback", GoogleCallback)

	log.Fatal(app.Listen(":8080"))
}
