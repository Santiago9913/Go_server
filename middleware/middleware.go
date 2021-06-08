package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/Santiago9913/Go_server/config"
)

func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config {
		Users: map[string]string{
			config.Config("USERNAME") : config.Config("PASSWORD")
		},
	}

	err := basicauth.New(cfg)
	return err
}