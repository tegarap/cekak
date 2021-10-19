package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tegarap/cekak/api/database"
	"github.com/tegarap/cekak/api/models"
	"github.com/tegarap/cekak/api/util"
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

	exist, _ := h.model.IsExist(input.ShortUrl)
	if exist != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(util.ResponseFail("Please Use Another Keyword", nil))
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

	url, err := h.model.Add(site)
	if err = c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(util.ResponseError("Fail to Create Short Url", nil))
	}

	sUrl := struct {
		ShortUrl string `json:"short_url"`
	}{
		ShortUrl: url.ShortUrl,
	}

	return c.Status(fiber.StatusCreated).
		JSON(util.ResponseSuccess("Success Create Short Url", sUrl))
}

func (h *UrlHandler) GetSiteHandler(c *fiber.Ctx) error {
	keyword := c.Params("keyword")

	url, err := h.model.GetSite(keyword)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(util.ResponseFail("Fail to Get Site Url", nil))
	}

	lUrl := struct {
		SiteUrl string `json:"site_url"`
	}{
		SiteUrl: url.LongUrl,
	}

	return c.Status(fiber.StatusOK).
		JSON(util.ResponseSuccess("Success Get Site Url", lUrl))
}

func (h *UrlHandler) DeleteHandler(c *fiber.Ctx) error {
	keyword := c.Params("keyword")

	if err := h.model.Delete(keyword); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(util.ResponseError("Fail to Delete Short Url", nil))
	}

	return c.Status(fiber.StatusOK).
		JSON(util.ResponseSuccess("Success Delete Short Url", nil))
}