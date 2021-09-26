package hint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/hint/:hint", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json; charset=utf-8")

		resp, err := http.Get(
			fmt.Sprintf(
				`http://clients1.google.com/complete/search?hl=en&client=firefox&output=toolbar&q=%s`,
				c.Params("hint"),
			),
		)

		if err != nil {
			return c.SendString("[]")
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return c.SendString("[]")
		}

		hints := string(body)
		hints = strings.Split(hints, "[")[2]
		hints = strings.Split(hints, "]")[0]

		return c.SendString(fmt.Sprintf("[%s]", hints))
	})
}
