package actions

import "github.com/gobuffalo/buffalo"

// AboutHandler is a default handler to serve up
// the about section on the home page.
func AboutHandler(c buffalo.Context) error {
	c.Set("title", "Home")
	return c.Render(200, r.HTML("index.html"))
}
