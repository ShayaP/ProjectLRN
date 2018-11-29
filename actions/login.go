package actions

import "github.com/gobuffalo/buffalo"

// LoginHandler is a default handler to serve up
// a login page.
func LoginHandler(c buffalo.Context) error {
	c.Set("title", "Login")
	return c.Render(200, r.HTML("login.html"))
}
