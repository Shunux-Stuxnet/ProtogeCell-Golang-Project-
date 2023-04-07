package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Shunux-Stuxnet/Project/config"
	"github.com/Shunux-Stuxnet/Project/initializers"
	"github.com/Shunux-Stuxnet/Project/models"
	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)

}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}
	var user models.User
	if err := json.Unmarshal(userData, &user); err != nil {
		panic(err)
	}
	initializers.InsertData(user)
	sess, err := store.Get(c)
	if err != nil {
		return c.SendString("Error in handeling the session")
	}
	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)
	if err := sess.Save(); err != nil {
		return c.SendString("Unable to set session for the user ")
	}
	return c.Redirect("/form")
}

func Index(c *fiber.Ctx) error {
	return c.Render("views/index.html", fiber.Map{})

}

func Report(c *fiber.Ctx) error {
	return c.Render("views/form.html", fiber.Map{})

}
