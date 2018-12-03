package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/cileonard/lrn/models"
)

func ResultHandler(c buffalo.Context) error {
	c.Set("title", "Results")
	filters := c.Session().Get("filter_info").(*Search)
	if filters != nil {
		name := filters.Name
		// lang := filters.Languages
		// loc := filters.Location
		// topics := filters.Topics
		tx := c.Value("tx").(*pop.Connection)
		u, err := models.GetUserByName(tx, name)
		if err != nil {
			return c.Error(401, err)
		}
		c.Session().Set("returned_user", u)
	}
	return c.Render(200, r.HTML("results.html"))
}