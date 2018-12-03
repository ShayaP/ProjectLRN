package actions
  
import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/cileonard/lrn/models"
	"fmt"
	)

// BrowseProfilesHandler is a default handler to serve up
// a home page.
func BrowseProfilesHandler(c buffalo.Context) error {
        c.Set("title", "Browse")
        // c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
        // c.Set("results", u)
        return c.Render(200, r.HTML("browseprofiles.html"))
}

func BrowseProfilesPOSTHandler(c buffalo.Context) error {
    c.Set("title", "Browse")
    s := &Search{}
	if err := c.Bind(s); err != nil {
		return c.Render(500, r.String(err.Error()))
	}
	fmt.Println(s)
    fmt.Println("LALALALALALALALALALALALAALALALALAAL")
    // return c.Render(200, r.JSON(s))
	name := s.Name
	fmt.Println(name)
	// lang := filters.Languages
	// loc := filters.Location
	// topics := filters.Topics
	tx := c.Value("tx").(*pop.Connection)
	u, err := models.GetUserByName(tx, name)
	if err != nil {
		return c.Error(401, err)
	}
	temp := []*models.User{u}
	c.Set("results", temp)
    // c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
    return c.Render(200, r.HTML("browseprofiles.html"))
}
