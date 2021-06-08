package router

import (
	"github.com/Santiago9913/Go_server/handler"
	"github.com/Santiago9913/Go_server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes (app *fiber.App) {
	api := app.Group("/api", logger.New(), middleware.AuthReq())


	//Routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}