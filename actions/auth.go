package actions

import (
	"fmt"
    "encoding/gob"

    "github.com/cileonard/lrn/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func init() {
	gothic.Store = App().SessionStore
    gob.Register(&models.User{})
	goth.UseProviders(
		google.New("190043352193-6gosbi41ard6f1itomqnd3u9kb831gtg.apps.googleusercontent.com", "mqkCA9p7dY1eej9TqmhkzFQx", fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
    u := &models.User{
        FirstName:  user.FirstName,
        LastName:   user.LastName,
        Email:      user.Email,
        GoogleID:   user.UserID,
    }
	c.Session().Set("userRequest", u)
    c.Session().Set("current_user", u.FirstName)
    //return c.Render(200, r.JSON(user))
	err = c.Session().Save()
	if err != nil {
		return c.Error(401, err)
	}
	// Do something with the user, maybe register them/sign them in
	c.Flash().Add("success", "Logged in!")
    return c.Redirect(302, "/auth/register")
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

func MoveUserObject(next buffalo.Handler) buffalo.Handler {
    fmt.Println("test")
    return func(c buffalo.Context) error{
        if userObj := c.Session().Get("userRequest"); userObj != nil {
            c.Set("userRequest", userObj)
            fmt.Println("WEALWJAN:EFJNEF")
        }
        return next(c)
    }
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Session().Get("current_user"); user != nil {
			name := user
			c.Set("name", name)
			c.Set("user", c.Session().Get("user"))
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
