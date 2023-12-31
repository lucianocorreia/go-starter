package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/pkg/context"
	"github.com/lucianocorreia/go-starter/pkg/handlers"
)

type errorHandler struct {
	handlers.Handler
}

func (e *errorHandler) Get(err error, ctx echo.Context) {
	if ctx.Response().Committed || context.IsCanceledError(err) {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code >= 500 {
		ctx.Logger().Error(err)
	} else {
		ctx.Logger().Info(err)
	}

	page := handlers.NewPage(ctx)
	page.Layout = "main"
	page.Title = http.StatusText(code)
	page.Name = "error"
	page.StatusCode = code
	// page.HTMX.Request.Enabled = false

	if err = e.RenderPage(ctx, page); err != nil {
		ctx.Logger().Error(err)
	}
}
