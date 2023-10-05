package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/pkg/handlers"
)

type (
	home struct {
		handlers.Handler
	}

	// post struct {
	// 	Title string
	// 	Body  string
	// }
)

func (h *home) Get(ctx echo.Context) error {
	page := handlers.NewPage(ctx)
	page.Layout = "main"
	page.Name = "home"
	page.Metatags.Description = "Welcom to the homepage."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	// page.Pager = controller.NewPager(ctx, 4)
	// page.Data

	return h.RenderPage(ctx, page)
}
