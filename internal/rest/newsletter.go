package rest

import (
	"github.com/gofiber/fiber/v2"
	domainNewsletter "github.com/liushiqi1001/go-whatsapp-web-multidevice/domains/newsletter"
	"github.com/liushiqi1001/go-whatsapp-web-multidevice/pkg/utils"
)

type Newsletter struct {
	Service domainNewsletter.INewsletterService
}

func InitRestNewsletter(app *fiber.App, service domainNewsletter.INewsletterService) Newsletter {
	rest := Newsletter{Service: service}
	app.Post("/newsletter/unfollow", rest.Unfollow)
	return rest
}

func (controller *Newsletter) Unfollow(c *fiber.Ctx) error {
	var request domainNewsletter.UnfollowRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	err = controller.Service.Unfollow(c.UserContext(), request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Status:  200,
		Code:    "SUCCESS",
		Message: "Success unfollow newsletter",
	})
}
