package main

import (
	"log"
	"profiley/database"
	"profiley/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := init_app()

	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes(app)
	profile_routes(app)

	app.Listen(":8000")
}

func init_app() error {
	err := load_env()

	if err != nil {
		panic(err)
	}

	err = load_db()

	if err != nil {
		panic(err)
	}

	return nil
}

func load_env() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't Load .env")
		return err
	}

	return nil
}

func load_db() error {
	err := database.ConnectMongo()

	if err != nil {
		log.Fatal("couldn't load the db")
		return err
	}

	return nil
}

func routes(app *fiber.App) {
	routes := app.Group("/")
	routes.Get("/health", handlers.CheckHealth)
	routes.Get("/:username", handlers.ValidateUsername)
}

func profile_routes(app *fiber.App) {
	profile_group := app.Group("/profile")
	profile_group.Get("/:username", handlers.GetProfile)
	profile_group.Post("/create", handlers.CreateProfile)
	profile_group.Post("/update/:username", handlers.UpdateProfile)
	profile_group.Post("/delete/:username", handlers.DeleteProfile)
}
