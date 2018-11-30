package actions

import (
    "fmt"

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
    fmt.Println("Here we are")
    u := &models.User{}
    if err := c.Bind(u); err != nil{
        fmt.Println("--------------")
        return c.Render(500, r.String(err.Error()))
    }
    userPrev := c.Session().Get("userObj").(*models.User)
    u.GoogleID = userPrev.GoogleID

    fmt.Println("~~~~~~~~")
    tx := c.Value("tx").(*pop.Connection)
    verrs, err := u.Create(tx)
    fmt.Println("////////////")
    if err != nil { //Error creating the account
        return c.Render(500, r.String(err.Error()))
        //return errors.WithStack(err)
    }

    fmt.Println("$$$$$$$$$$$")
    if verrs.HasAny() {
        c.Set("user", u)
        c.Set("errors", verrs)
	    c.Set("title", "Register - Error")
        return c.Render(200, r.HTML("register.html"))
    }

    c.Session().Set("user", u)
    return c.Redirect(302, "/")
}
