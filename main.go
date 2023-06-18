package main

import (
	"log"

	"github.com/MSyahidAlFajri/websocket-heroku/module"
	"github.com/MSyahidAlFajri/websocket-heroku/url"
	"github.com/aiteung/musik"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
