package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"profiley/database"
	"profiley/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateProfile(c *fiber.Ctx) error {
	profile := new(models.Profile)

	if err := c.BodyParser(&profile); err != nil {
		fmt.Println("error: ", err)
		return c.SendStatus(200)
	}

	collection := database.GetCollection("profiles")

	var posts []models.Profile
	cursor, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var elem models.Profile
		err := cursor.Decode(&elem)

		if err != nil {
			panic(err)
		}

		posts = append(posts, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	for _, post := range posts {
		if profile.Username == post.Username {
			return c.SendString("Sorry, this username already exists. Please change it.")
		} else {
			_, err := collection.InsertOne(context.TODO(), profile)
			fmt.Println("created profile")

			if err != nil {
				panic(err)
			}
		}
	}

	return c.SendString("success")
}

func GetProfile(c *fiber.Ctx) error {
	collection := database.GetCollection("profiles")
	param_username := c.Params("username")
	fmt.Println(param_username)

	profile_filter := bson.D{{Key: "username", Value: param_username}}
	fmt.Println(profile_filter)

	cur, err := collection.Find(context.TODO(), profile_filter)

	if err != nil {
		panic(err)
	}

	var results []models.Profile
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	return c.SendString("success")
}
