package actions

import (
	"fmt"

	//"github.com/pkg/errors"

	"github.com/cileonard/lrn/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// RegisterRegisterHandler default implementation.
func RegisterHandler(c buffalo.Context) error {
	c.Set("title", "Register")
	c.Set("NoShow", true)

	fmt.Println("UUUUUUUUUUUUUUUUUUUU")
	fmt.Println(c.Session().Get("userID"))
	fmt.Println("FFFFFFFFFFFFFFFFFFff")
	return c.Render(200, r.HTML("register.html"))
}

func RegisterPOSTHandler(c buffalo.Context) error {
	c.Set("NoShow", true)
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return c.Render(500, r.String(err.Error()))
	}
	//userPrev := c.Session().Get("user").(*models.User)
	u.GoogleID = c.Session().Get("userID").(string)

    if u.Gender != 3 {
        u.OtherSpecify = "-"
    }

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil { //Error creating the account
		return c.Render(500, r.String(err.Error()))
		//return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("userRequest", u)
		c.Set("errors", verrs)
		c.Set("title", "Register - Error")
		return c.Render(200, r.HTML("register.html"))
	}

	c.Session().Set("current_user", u.FirstName)
	c.Session().Set("user", u)
	return c.Redirect(302, "/")
}
