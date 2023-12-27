package handlers

import (
	"context"
	"fmt"
	"profiley/database"
	"profiley/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(ctx *fiber.Ctx) error {
	profile := new(models.Profile)

	if err := ctx.BodyParser(&profile); err != nil {
		fmt.Println("error: ", err)
		return ctx.SendStatus(200)
	}

	collection := database.GetCollection("profiles")
	result, err := collection.InsertOne(context.TODO(), profile)

	fmt.Println(result)

	if err != nil {
		panic(err)
	}

	return ctx.SendString("success")
}
