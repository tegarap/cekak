package routes

import (
	"github.com/tegarap/cekak/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.UrlHandler) {
	app.Post("", h.CreateShortUrlHandler)
	app.Get("", h.GetAllSiteHandler)
	app.Get("/:keyword", h.GetSiteHandler)
	app.Delete("/:keyword", h.DeleteHandler)
	app.Get("/redirect/:keyword", h.RedirectHandler)
}