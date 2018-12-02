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
    u, err := models.GetUserByGID(tx, gu.UserID)

	if err != nil {
        u := &models.User{}
        u.FirstName = gu.FirstName
        u.LastName = gu.LastName
        u.GoogleID = gu.UserID
        u.Email = gu.Email
        c.Session().Set("userID", gu.UserID)
		c.Session().Set("userRequest", u)
		return c.Redirect(302, "/auth/register")
	} else {
		c.Session().Set("user", u)
		err = c.Session().Save()
		if err != nil {
			return c.Error(401, err)
		}
		// Do something with the user, maybe register them/sign them in
		return c.Redirect(302, "/")
	}
}

// func AuthDestroy(c buffalo.Context) error {
// 	c.Session().Clear()
// 	err := c.Session().Save()
// 	if err != nil {
// 		return c.Error(401, err)
// 	}
// 	c.Flash().Add("success", "Logged out!")
// 	return c.Redirect(302, "/")
// }

func MoveUserObject(next buffalo.Handler) buffalo.Handler {
    fmt.Println("test")
    return func(c buffalo.Context) error{
        if userObj := c.Session().Get("userRequest"); userObj != nil {
            u := userObj.(*models.User)
            c.Set("userRequest", u)
            //c.Set("userID", u.GoogleID)
            //fmt.Println(u.GoogleID)
        }
        return next(c)
    }
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Session().Get("user"); user != nil {
			c.Set("user", user)
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Session().Get("user"); user == nil {
			fmt.Println("GONNA REDIRECT YOU")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
