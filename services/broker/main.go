package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func getEnv(v string) string {
	return fmt.Sprintf("%s = \"%s\"\n", v, os.Getenv(v))
}
func getHeader(c *fiber.Ctx, h string) string {
	return fmt.Sprintf("%s = \"%s\"\n", h, c.Get(h))
}

func main() {
	app := fiber.New()

	port := os.Getenv("SERVICE_PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		log.Print(".")
		msg := "Broker:\nEnvs:\n"
		msg += getEnv("GAE_SERVICE")
		msg += getEnv("SERVICE_PORT")
		msg += getEnv("PORT")

		// read headers
		msg += "\nHeaders:\n"
		msg += getHeader(c, "X-Appengine-User-Ip")
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":" + port))
}
