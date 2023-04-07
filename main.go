package main

import (
	"github.com/Shunux-Stuxnet/Project/config"
	"github.com/Shunux-Stuxnet/Project/initializers"
	"github.com/Shunux-Stuxnet/Project/server"
)

// var app *fiber.App

// func init() {
// 	initializers.ConnectDB()
//
// 	app = fiber.New()
//
// }

func main() {
	// app := fiber.New()
	config.GoogleConfig()
	initializers.ConnectDB()
	//app.Get("/protected", server.middleware, func(c *fiber.Ctx) error {
	//	return c.SendString("You are authenticated!")
	//})
	//
	// app.Get("/google_login", controllers.GoogleLogin)
	//
	// app.Get("/google_callback", controllers.GoogleCallback)
	//
	// app.Listen(":8080")
	server.Serve()

}
