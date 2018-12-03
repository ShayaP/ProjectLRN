package actions
  
import "github.com/gobuffalo/buffalo"

// BrowseProfilesHandler is a default handler to serve up
// a home page.
func BrowseProfilesHandler(c buffalo.Context) error {
        c.Set("title", "Browse")
        c.Set("results", []int{1,2,3,4,5})
        return c.Render(200, r.HTML("browseprofiles.html"))
    }