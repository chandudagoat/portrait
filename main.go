package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	load_env()
}

func load_env() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't Load .env")
	}
}
