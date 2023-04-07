package server

import "github.com/gofiber/fiber/v2/middleware/session"

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)
