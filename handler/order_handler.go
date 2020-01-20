package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func OrderHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go

	r := c.Request()
	return c.Render(http.StatusOK, "order.html", map[string]interface{}{
		"name":   "Order",
		"id":     r.URL.Query()["id"][0],
		"nama":   r.URL.Query()["nama"][0],
		"gambar": r.URL.Query()["gambar"][0],
	})
}
