package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"profiley/database"
	"profiley/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUsername(c *fiber.Ctx) error {
	collection := database.GetCollection("profiles")
	param_username := c.Params("username")

	profile_filter := bson.D{{Key: "username", Value: param_username}}

	cur, err := collection.Find(context.TODO(), profile_filter)

	if err != nil {
		panic(err)
	}

	var results []models.Profile
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return c.JSON(results)
}

func CreateProfile(c *fiber.Ctx) error {
	profile := new(models.Profile)

	if err := c.BodyParser(&profile); err != nil {
		fmt.Println("error: ", err)
		c.SendStatus(200)
	}

	collection := database.GetCollection("profiles")
	collection.InsertOne(context.TODO(), profile)

	fmt.Printf("created profile with username %s", profile.Username)

	return c.SendString("success")
}

// func CreateProfile(c *fiber.Ctx) error {
// 	profile := new(models.Profile)

// 	if err := c.BodyParser(&profile); err != nil {
// 		fmt.Println("error: ", err)
// 		return c.SendStatus(200)
// 	}

// 	collection := database.GetCollection("profiles")

// 	var posts []models.Profile
// 	cursor, err := collection.Find(context.TODO(), bson.D{{}})

// 	if err != nil {
// 		panic(err)
// 	}

// 	for cursor.Next(context.TODO()) {
// 		var elem models.Profile
// 		err := cursor.Decode(&elem)

// 		if err != nil {
// 			panic(err)
// 		}

// 		posts = append(posts, elem)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	cursor.Close(context.TODO())

// 	for _, post := range posts {
// 		if profile.Username == post.Username {
// 			return c.SendString("Sorry, this username already exists. Please change it.")
// 		} else {
// 			_, err := collection.InsertOne(context.TODO(), profile)
// 			fmt.Println("created profile")

// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}

// 	return c.SendString("success")
// }

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
		return c.SendString(string(res))
	}

	return c.SendString("success")
}

func UpdateProfile(c *fiber.Ctx) error {
	collection := database.GetCollection("profiles")
	param_username := c.Params("username")

	profile_filter := bson.D{{Key: "username", Value: param_username}}
	filter := new(models.Profile)

	if err := c.BodyParser(&filter); err != nil {
		panic(err)
	}

	updated_filter := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: filter.Name},
			{Key: "pronouns", Value: filter.Pronouns},
			{Key: "bio", Value: filter.Bio},
			{Key: "links", Value: filter.Links},
		},
	}}

	_, err := collection.UpdateOne(context.TODO(), profile_filter, updated_filter)

	if err != nil {
		panic(err)
	}

	return c.SendString("success")
}

func DeleteProfile(c *fiber.Ctx) error {
	collection := database.GetCollection("profiles")
	username := c.Params("username")

	filter := bson.D{{Key: "username", Value: username}}

	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	success_msg := "successfully deleted user" + username

	return c.SendString(success_msg)
}
