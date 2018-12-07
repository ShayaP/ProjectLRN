package actions

import "github.com/gobuffalo/buffalo"

// FindHandler is a default handler to serve up
// a home page.
func FindHandler(c buffalo.Context) error {
	c.Set("title", "Find")
	return c.Render(200, r.HTML("find.html"))
}

