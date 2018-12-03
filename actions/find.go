package actions

import "github.com/gobuffalo/buffalo"

type Search struct {
	Name string
	Languages string
	Location string
	Topics string
}

// FindHandler is a default handler to serve up
// a home page.
func FindHandler(c buffalo.Context) error {
	c.Set("title", "Find")
	return c.Render(200, r.HTML("find.html"))
}

func FindPOSTHandler(c buffalo.Context) error {
	c.Set("NoShow", true)
	s := &Search{}
	if err := c.Bind(s); err != nil {
		return c.Render(500, r.String(err.Error()))
	}
	c.Session().Set("filter_info", s)
	return c.Redirect(302, "/results")
}
