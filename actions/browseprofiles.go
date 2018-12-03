package actions
  
import "github.com/gobuffalo/buffalo"

// BrowseProfilesHandler is a default handler to serve up
// a home page.
func BrowseProfilesHandler(c buffalo.Context) error {
        c.Set("title", "Browse")
        return c.Render(200, r.HTML("browseprofiles.html"))
    }