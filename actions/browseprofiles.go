package actions
  
import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/cileonard/lrn/models"
	"fmt"
	"strconv"
	)

type Search struct {
	Name string
	Languages string
	Location int
	Topics string
}

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
    // return c.Render(200, r.JSON(s))
	name := s.Name
	fmt.Println(name)
	tx := c.Value("tx").(*pop.Connection)

	if name == "" {
		lang := s.Languages
		loc := strconv.Itoa(s.Location)
		topics := s.Topics
		list, err := models.GetFilteredUsers(tx, loc, lang, topics)
		if err != nil {
			return c.Error(401, err)
		}
		fmt.Println()
		fmt.Println(list)
		users := []*models.User{} 
		for _, userinfo := range list {
			u, err := models.GetUserByGID(tx, userinfo.GoogleID)
			if err != nil {
				return c.Error(401, err)
			}
			users = append(users, u)
		}
		c.Set("results", users)
	} else {
		u, err := models.GetUserByName(tx, name)
		if err != nil {
			return c.Error(401, err)
		}
		temp := []*models.User{u}
		c.Set("results", temp)
	}
    // c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
    return c.Render(200, r.HTML("browseprofiles.html"))
}
