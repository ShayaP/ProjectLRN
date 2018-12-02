package actions

import (
	"fmt"
    "encoding/gob"

    "github.com/cileonard/lrn/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/gobuffalo/pop"
)

func init() {
	gothic.Store = App().SessionStore
    gob.Register(&models.User{})
	goth.UseProviders(
		google.New("190043352193-6gosbi41ard6f1itomqnd3u9kb831gtg.apps.googleusercontent.com", "mqkCA9p7dY1eej9TqmhkzFQx", fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx := c.Value("tx").(*pop.Connection)
	u := &models.User{}
	err = tx.Find(u, gu.UserID)
	fmt.Println(err.Error())
	if err != nil {
		return c.Redirect(302, "/register")
	} else {
		c.Session().Set("userObj", u)
		
		err = c.Session().Save()
		if err != nil {
			return c.Error(401, err)
		}
		// Do something with the user, maybe register them/sign them in
		return c.Redirect(302, "/")
	}
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
            fmt.Println("AAAAAAAAAAAAAAAAAAAAaaaa")
		}
		return next(c)
	}
}
