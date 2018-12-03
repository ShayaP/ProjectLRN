package actions

import (
	"github.com/gobuffalo/buffalo"
	// "github.com/cileonard/lrn/models"
	// "github.com/gobuffalo/pop"
)

// FindHandler is a default handler to serve up
// a home page.
func RequestPageHandler(c buffalo.Context) error {
	c.Set("title", "RequestPage")
	// user := c.Session().Get("user")
	
	return c.Render(200, r.HTML("requestpage.html"))
}
