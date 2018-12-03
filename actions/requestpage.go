package actions

import "github.com/gobuffalo/buffalo"

// FindHandler is a default handler to serve up
// a home page.
func RequestPageHandler(c buffalo.Context) error {
	c.Set("title", "RequestPage")
	c.Set("sentRequests", GetRequestsSent(c))
	return c.Render(200, r.HTML("requestpage.html"))
}

// Get the requests the user sent
func GetRequestsSent(c buffalo.Context) [][]string {
	// psuedo request data
	requests := [][]string {
		{"Bobby Tefla", "5" ,"Zoology", "Pending", "4/2/18"},
		{"Josh Ben", "4" ,"U.S. History", "Rejected", "6/2/18"},
		{"Hannah Benson", "4" ,"Linear ALgebra", "Accepted", "7/2/18"},
		{"Lizzy Queen", "5" ,"Introduction to Chemistry", "Accepted", "10/2/18"},
	}
	return requests
}
