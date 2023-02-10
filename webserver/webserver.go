package webserver

import (
	"fmt"
	"main/discord"
	"main/webserver/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
	"main/utils"
)

type WebServer struct {
	App 	*fiber.App
	Engine  *handlebars.Engine

	Discord *discord.Discord
}


/* Functions */

func onRequestMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func NewWebServer(token string) *WebServer {
	engine := handlebars.New("webserver/static/html/", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine.Reload(true),
	})

	app.Static("/static", "webserver/static")
	app.Use(onRequestMiddleware)

	app.Get("/", handlers.IndexHandler)
	app.Get("/:user", handlers.ProfileHandler)


	utils.Discord = discord.New(token)
	return &WebServer{
		App: app,
		Engine: engine,
		Discord: utils.Discord,
	}
}

func (ws *WebServer) Start(port int) error {
	fmt.Printf("Starting webserver: http://localhost:%d\n", port)
	return ws.App.Listen(fmt.Sprintf(":%d", port))
}
