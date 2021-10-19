package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarap/cekak/api/handlers"
)

func Routes(app *fiber.App, h *handlers.UrlHandler) {
	app.Post("/create", h.CreateShortUrlHandler)
	app.Get("/all", h.GetAllSiteHandler)
	app.Get("/:keyword", h.GetSiteHandler)
}