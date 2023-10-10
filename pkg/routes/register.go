package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/pkg/context"
	"github.com/lucianocorreia/go-starter/pkg/handlers"
)

type (
	register struct {
		handlers.Handler
	}

	registerForm struct {
		name string `form:"name" validate:"required"`
	}
)

func (c *register) Get(ctx echo.Context) error {
	page := handlers.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "register"
	page.Title = "Register"
	page.Form = registerForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*registerForm)
	}

	return c.RenderPage(ctx, page)
}
