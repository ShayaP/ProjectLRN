package actions

import (
    //"fmt"

	"github.com/cileonard/lrn/models"

    "github.com/gobuffalo/uuid"
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

type responseNewRequest struct{
    ReceiverID  string  `json:"receiver"`
    Topic       string  `json:"topic"`
}

func RequestPagePOSTHandler(c buffalo.Context) error {
    user := c.Session().Get("user").(*models.User)

    var resp map[string][]string
    resp = c.Request().Form

	tx := c.Value("tx").(*pop.Connection)
    sendID := user.ID
    recID, err := uuid.FromString(resp["receiver"][0])
    if err != nil {
        return c.Render(401, r.String(err.Error()))
    }
    req := &models.Request{
        Status:     2,
        SenderID:   sendID,
        ReceiverID: recID,
        Topic:      resp["topic"][0],
    }

    verrs, err := tx.ValidateAndCreate(req)

    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }
    if verrs.HasAny(){
        return c.Render(500, r.String("Validation errors"))
    }

    return c.Render(200, r.String("Created new request"))

	//if err := c.Bind(u); err != nil {
	//	return c.Render(500, r.String(err.Error()))
	//}

	//sender  = c.Session().Get("user").(string)
	//reciever  =  	
	//topic = 

	//verrs, err := request.CreateRequest(tx, sender, reciever, topic)

	//if err != nil { //Error creating the request
	//	return c.Render(500,  r.String(err.Error()))
	//}

}

func UpdateRequestSentPOSTHandler(c buffalo.Context) error {
    var resp map[string][]string
    resp = c.Request().Form

	tx := c.Value("tx").(*pop.Connection)

    reqID, err := uuid.FromString(resp["request"][0])
    if err != nil {
        return c.Render(401, r.String(err.Error()))
    }

    req := &models.Request{}
    err = tx.Find(req, reqID)

    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }

    req.Status = 3


    verrs, err := tx.ValidateAndUpdate(req)

    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }
    if verrs.HasAny(){
        return c.Render(500, r.String("Validation errors"))
    }
    return c.Render(200, r.String("Created new request"))
}
func UpdateRequestReceivedPOSTHandler(c buffalo.Context) error {
    var resp map[string][]string
    resp = c.Request().Form

	tx := c.Value("tx").(*pop.Connection)

    reqID, err := uuid.FromString(resp["request"][0])
    if err != nil {
        return c.Render(401, r.String(err.Error()))
    }

    req := &models.Request{}
    err = tx.Find(req, reqID)

    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }

    decision := resp["decision"][0]
    if (decision == "1"){
        req.Status = 1
    }else{
        req.Status = 4
    }


    verrs, err := tx.ValidateAndUpdate(req)

    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }
    if verrs.HasAny(){
        return c.Render(500, r.String("Validation errors"))
    }
    return c.Render(200, r.String("Created new request"))
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
