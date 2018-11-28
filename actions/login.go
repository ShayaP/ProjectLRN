package actions

import "github.com/gobuffalo/buffalo"

// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
	c.Set("title", "Login")
	return c.Render(200, r.HTML("login.html"))
}
