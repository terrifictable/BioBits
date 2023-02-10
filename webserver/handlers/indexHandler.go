package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})

	//t := template.Must(template.ParseGlob("webserver/static/html/*"))
	//err := t.ExecuteTemplate(w, "index.hbs", nil)
	//if err != nil {
	//	fmt.Fprintf(w, "Error executing template: %s\n", err.Error())
	//}
}
