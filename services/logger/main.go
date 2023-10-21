package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("SERVICE_PORT")
	log.Printf("Listening on: %s ...", port)
	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello from %s listening on: %s",
			os.Getenv("GAE_SERVICE"),
			os.Getenv("SERVICE_PORT"))
		log.Println(msg)
		return c.SendString("logger\n" + msg + "\nREADY >_")
	})

	log.Fatal(app.Listen(":" + port))
}
