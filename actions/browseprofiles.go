package actions
  
import "github.com/gobuffalo/buffalo"

// BrowseProfilesHandler is a default handler to serve up
// a home page.
func BrowseProfilesHandler(c buffalo.Context) error {
        c.Set("title", "Browse")
        c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
        return c.Render(200, r.HTML("browseprofiles.html"))
}

func BrowseProfilesPOSTHandler(c buffalo.Context) error {
    c.Set("title", "Browse")
    c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
    return c.Render(200, r.HTML("browseprofiles.html"))
}
