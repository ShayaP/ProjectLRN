package actions

import (
	"github.com/gobuffalo/buffalo"
)

// RegisterRegisterHandler default implementation.
func RegisterHandler(c buffalo.Context) error {
	c.Set("title", "Register")
	return c.Render(200, r.HTML("register.html"))
}
