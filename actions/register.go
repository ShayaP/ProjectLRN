package actions

import (
    //"fmt"

    //"github.com/pkg/errors"

    "github.com/cileonard/lrn/models"

    "github.com/gobuffalo/pop"
	"github.com/gobuffalo/buffalo"
)

// RegisterRegisterHandler default implementation.
func RegisterHandler(c buffalo.Context) error {
	c.Set("title", "Register")
	return c.Render(200, r.HTML("register.html"))
}

func RegisterPOSTHandler(c buffalo.Context) error {
    u := &models.User{}
    if err := c.Bind(u); err != nil{
        return c.Render(500, r.String(err.Error()))
    }
    userPrev := c.Session().Get("userObj").(*models.User)
    u.GoogleID = userPrev.GoogleID

    tx := c.Value("tx").(*pop.Connection)
    verrs, err := u.Create(tx)
    if err != nil { //Error creating the account
        return c.Render(500, r.String(err.Error()))
        //return errors.WithStack(err)
    }

    if verrs.HasAny() {
        c.Set("user", u)
        c.Set("errors", verrs)
	    c.Set("title", "Register - Error")
        return c.Render(200, r.HTML("register.html"))
    }

    c.Session().Set("user", u)
    return c.Redirect(302, "/")
}
