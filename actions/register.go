package actions

import (
    "github.com/gobuffalo/buffalo"
    "fmt"
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
