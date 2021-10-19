package handlers

import (
	"github.com/tegarap/cekak/api/database"
	"github.com/tegarap/cekak/api/models"
	"github.com/tegarap/cekak/api/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UrlHandler struct {
	model database.SiteModelImpl
}

func NewUrlHandler(model database.SiteModelImpl) *UrlHandler {
	return &UrlHandler{model: model}
}

func (h *UrlHandler) GetAllSiteHandler(c *fiber.Ctx) error {
	allUrl, err := h.model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(util.ResponseError("Cannot Get All Site", nil))
	}
	return c.Status(fiber.StatusOK).
		JSON(util.ResponseSuccess("Success Get All Site", allUrl))
}

func (h *UrlHandler) CreateShortUrlHandler(c *fiber.Ctx) error {
	var err error

	input := new(models.Site)
	err = c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(util.ResponseError("Server Error", nil))
	}

	//input.LongUrl = strings.Replace(strings.ToLower(input.LongUrl), "https://", "http://", 1)

	validate := validator.New()
	if err = validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(util.ResponseFail("Invalid Site Url", nil))
	}

	site := models.Site{
		LongUrl:  input.LongUrl,
		ShortUrl: input.ShortUrl,
	}

	Url, err := h.model.Add(site)
	if err = c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(util.ResponseError("Fail to Create Shortlink", nil))
	}

	sLink := struct {
		Site string `json:"site"`
	}{
		Site: Url.ShortUrl,
	}

	return c.Status(fiber.StatusCreated).
		JSON(util.ResponseSuccess("Success Create Shortlink", sLink))
}

func (h *UrlHandler) GetSiteHandler(c *fiber.Ctx) error {
	//	keyword := c.Params("keyword")
	//
	//	var site models.Site
	//
	//
	//
	return c.Status(fiber.StatusOK).JSON(util.ResponseSuccess("Forward to Link", nil))
}
