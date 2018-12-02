package actions

import "github.com/gobuffalo/buffalo"

// FindHandler is a default handler to serve up
// a home page.
func LogoutHandler(c buffalo.Context) error {
	// c.Set("title", "Find")
	c.Session().Clear()
	err := c.Session().Save()
	if err != nil {
		return c.Error(401, err)
	}
	// c.Flash().Add("success", "Logged out!")
	return c.Redirect(302, "/")
}
