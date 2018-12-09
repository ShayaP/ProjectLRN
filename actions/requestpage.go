package actions

import (
    //"fmt"

	"github.com/cileonard/lrn/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// FindHandler is a default handler to serve up
// a home page.
func RequestPageHandler(c buffalo.Context) error {
    user := c.Session().Get("user").(*models.User)

	tx := c.Value("tx").(*pop.Connection)
    allSentRequests, err := models.GetRequestsSent(user, tx)
    if err != nil {
        return c.Render(500, r.String(err.Error()))
    }
    allReceivedRequests, err := models.GetRequestsReceived(user, tx)
    if err != nil {
        return c.Render(500, r.String(err.Error()))
    }


    ParsedSentRequests, err := SplitRequestsByStatus(*allSentRequests, true, tx)
    if err != nil {
        return c.Render(500, r.String(err.Error()))
    }
    ParsedRecievedRequests, err := SplitRequestsByStatus(*allReceivedRequests, false, tx)
    if err != nil {
        return c.Render(500, r.String(err.Error()))
    }

	c.Set("title", "RequestPage")
	c.Set("sentRequests", ParsedSentRequests)
	c.Set("receivedRequests", ParsedRecievedRequests)
	return c.Render(200, r.HTML("requestpage.html"))
}


//func RequestPagePOSTHandler(c buffalo.Context) error {
//}

// Get the requests the user sent
func GetRequestsSent(c buffalo.Context) [][]string {
	// psuedo request data
	requests := [][]string {
		{"Bobby Tefla", "5" ,"Zoology", "Pending", "4/2/18"},
		{"Josh Ben", "4" ,"U.S. History", "Rejected", "6/2/18"},
		{"Hannah Benson", "4" ,"Linear Algebra", "Accepted", "7/2/18"},
		{"Lizzy Queen", "5" ,"Introduction to Chemistry", "Accepted", "10/2/18"},
	}
	return requests
}

// Get the requests the user recieved
func GetRequestsReceived(c buffalo.Context) [][]string {
	// psuedo request data
	requests := [][]string {
		{"Benny John", "3" ,"Zoology", "Pending", "9/2/18"},
		{"Eliza Lu", "1" ,"U.S. History", "Pending", "6/2/18"},
		{"Paddy Pie", "4" ,"Linear Algebra", "Rejected", "7/2/16"},
		{"Minnie Jane", "5" ,"Introduction to Chemistry", "Rejected", "10/21/18"},
	}
	return requests
}
