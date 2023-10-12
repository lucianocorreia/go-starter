package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/lucianocorreia/go-starter/pkg/flash"
)

// Page consists of all data that will be used to render a page response for a given controller.
type Page struct {
	AppName string
	Title   string

	// Context stores the request context
	Context echo.Context

	// ToURL is a function to convert a route name and optional route parameters to a URL
	ToURL func(name string, params ...interface{}) string

	Path       string
	URL        string
	Form       interface{}
	Layout     string
	Name       string
	IsHome     bool
	IsAuth     bool
	StatusCode int
	// Pager

	// User is the authenticated user
	// User ...

	// Metatags stores metatag values
	Metatags struct {
		Description string
		Keywords    []string
	}

	CSRF string

	Headers   map[string]string
	RequestID string
}

// NewPage creates a new page
func NewPage(ctx echo.Context) Page {
	p := Page{
		Context:    ctx,
		ToURL:      ctx.Echo().Reverse,
		Path:       ctx.Request().URL.Path,
		URL:        ctx.Request().URL.String(),
		StatusCode: http.StatusOK,
		// Pager:      NewPager(ctx, DefaultItemsPerPage),
		Headers:   make(map[string]string),
		RequestID: ctx.Response().Header().Get(echo.HeaderXRequestID),
	}

	p.IsHome = p.Path == "/"

	if csrf := ctx.Get(echomw.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	// if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
	// 	p.IsAuth = true
	// 	p.AuthUser = u.(*ent.User)
	// }

	// p.HTMX.Request = htmx.GetRequest(ctx)

	return p
}

// GetMessages gets all flash messages for a given type.
// This allows for easy access to flash messages from the templates.
func (p Page) GetMessages(typ flash.FlashType) []template.HTML {
	strs := flash.Get(p.Context, typ)
	ret := make([]template.HTML, len(strs))
	for k, v := range strs {
		ret[k] = template.HTML(v)
	}
	return ret
}
