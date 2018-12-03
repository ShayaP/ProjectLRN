package actions

import "github.com/gobuffalo/buffalo"

func FindHandler(c buffalo.Context) error {
	c.Set("title", "Results")
	filters := c.Session().Get("filter_info")
	if filters != nil {
		name := filters.Name
		lang := filters.Laguages
		loc := filters.Location
		topics := filters.Topics
		
	}
	return c.Render(200, r.HTML("results.html"))
}