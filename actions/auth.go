package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		google.New("190043352193-6gosbi41ard6f1itomqnd3u9kb831gtg.apps.googleusercontent.com", "mqkCA9p7dY1eej9TqmhkzFQx", fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	c.Session().Set("current_user", user.Name)
	c.Session().Set("userObj", user)
    err = c.Session().Save()
	fmt.Println("======")
	fmt.Println(c.Session().Get("userObj"))
    fmt.Println(user)
	fmt.Println("~~~~~~")
	if err != nil {
		return c.Error(401, err)
	}
	// Do something with the user, maybe register them/sign them in
	c.Flash().Add("success", "Logged in!")
	return c.Redirect(302, "/register")
}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	err := c.Session().Save()
	if err != nil {
		return c.Error(401, err)
	}
	c.Flash().Add("success", "Logged out!")
	return c.Redirect(302, "/")
}
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Session().Get("current_user"); user != nil {
			name := user
			c.Set("name", name)
            c.Set("userObj", c.Session().Get("userObj"))
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Session().Get("current_user"); user == nil {
			c.Flash().Add("danger", "You must be logged in to see that page!")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
