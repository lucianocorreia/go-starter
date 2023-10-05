package routes

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/lucianocorreia/go-starter/config"
	"github.com/lucianocorreia/go-starter/pkg/middleware"
	"github.com/lucianocorreia/go-starter/pkg/services"
)

// BBuildRouter builds the router
func BuildRouter(c *services.Container) {
	// Static files
	// TODO: add cache control
	c.Web.Group("").Static(config.StaticPrefix, config.StaticDir)

	// Non static files route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{RedirectCode: http.StatusMovedPermanently}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.Logger(),
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		// middleware.LoadAuthenticatedUser(c.Auth),
		// middleware.ServeCachedPage(c.Cache),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

}
