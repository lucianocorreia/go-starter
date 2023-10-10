package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/pkg/context"
)

// RequireNoAuthentication requires that the user not be authenticated in order to proceed
func RequireNoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u != nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
