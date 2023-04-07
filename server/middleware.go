package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddeware(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/google") {
		if sess, _ := store.Get(c); sess.Get(AUTH_KEY) != nil {
			return c.Redirect("/form")
		}

		return c.Next()
	}
	sess, err := store.Get(c)
	if err != nil {
		return c.Redirect("/google_login")
	}

	if sess.Get(AUTH_KEY) != nil {
		return c.Next()
	}

	return c.Redirect("/google_login")
}
