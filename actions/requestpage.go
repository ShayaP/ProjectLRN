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


func RequestPagePOSTHandler(c buffalo.Context) error {	
    user := c.Session().Get("user").(*models.User)

    return c.Render(200, r.JSON(user))
    //return c.Render(200, r.JSON(c.Request().Form))

	//request := &models.Request{}

	//if err := c.Bind(u); err != nil {
	//	return c.Render(500, r.String(err.Error()))
	//}

	//sender  = c.Session().Get("user").(string)
	//reciever  =  	
	//topic = 

	//tx := c.Value("tx").(*pop.Connection)	
	//verrs, err := request.CreateRequest(tx, sender, reciever, topic)

	//if err != nil { //Error creating the request
	//	return c.Render(500,  r.String(err.Error()))
	//}

}

/**
func acceptRequest(c buffalo.Context) error {
	user := c.Session().Get("user").(*models.User)
	
	requests := GetRequestRecieved(c)

	allReceivedRequests, err := models.GetRequestsReceived(user, tx)
    	if err != nil {
        	return c.Render(500, r.String(err.Error()))
    	}


}

func declineRequest(c Buffalo.context) error {


}*/
