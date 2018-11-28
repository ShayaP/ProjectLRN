package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

// RegisterRegisterHandler default implementation.
func RegisterHandler(c buffalo.Context) error {
	//if(c.Session().Get("userObj") == nil){
	//    fmt.Println("WHATTHEFUCK")
	//}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(c.Session().Get("userObj"))
	c.Set("title", "Register")
	return c.Render(200, r.HTML("register.html"))
}
