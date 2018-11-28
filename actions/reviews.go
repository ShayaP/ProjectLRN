package actions

import "github.com/gobuffalo/buffalo"

// FindHandler is a default handler to serve up
// a home page.
func ReviewHandler(c buffalo.Context) error {
	c.Set("title", "RequestPage")
	return c.Render(200, r.HTML("reviewspage.html"))
}
