package routes

import (
	"github.com/tegarap/cekak/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.UrlHandler) {
	app.Post("/create", h.CreateShortUrlHandler)
	app.Get("/all", h.GetAllSiteHandler)
	app.Get("/:keyword", h.GetSiteHandler)
}