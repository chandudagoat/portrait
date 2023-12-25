package handlers

import (
	"fmt"
	"profiley/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(ctx *fiber.Ctx) error {
	profile := new(models.Profile)

	if err := ctx.BodyParser(&profile); err != nil {
		fmt.Println("error: ", err)
		return ctx.SendStatus(200)
	}

	fmt.Println("name: ", profile.Name)
	fmt.Println("pronouns: ", profile.Pronouns)
	fmt.Println("bio: ", profile.Bio)
	fmt.Println("links: ", profile.Links)

	return ctx.SendString("success")
}
