package actions

import "github.com/gobuffalo/buffalo"

// UpdateProfileHandler is a default handler to serve up
// an editable profile page.
func UpdateProfileHandler(c buffalo.Context) error {
	c.Set("title", "Update Profile")
	return c.Render(200, r.HTML("update-profile.html"))
}
