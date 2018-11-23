package actions

// import (
// 	"fmt"

// 	"github.com/gobuffalo/buffalo"
// 	"github.com/gobuffalo/envy"
// 	"github.com/markbates/goth"
// 	"github.com/markbates/goth/gothic"
// 	"github.com/markbates/goth/providers/google"
// )

// func init() {
// 	gothic.Store = App().SessionStore

// 	goth.UseProviders(
// 		google.New(envy.Get("CLIENT_ID", ""), envy.Get("SECRET_ID", ""), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
// 	)
// }

// func AuthCallback(c buffalo.Context) error {
// 	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
// 	if err != nil {
// 		return c.Error(401, err)
// 	}

// 	//Do user DB things
// 	return c.Render(200, r.JSON(user))
// }
