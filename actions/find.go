package actions

import "github.com/gobuffalo/buffalo"

// FindHandler is a default handler to serve up
// a home page.
func FindHandler(c buffalo.Context) error {
	languages := uinfo.GetLanguages()
	c.Set("title", "Find")
	c.Set("languages", languages)
	return c.Render(200, r.HTML("find.html"))
}

