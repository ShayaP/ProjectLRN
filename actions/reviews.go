package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/cileonard/lrn/models"
	"github.com/gobuffalo/pop"
)


// ReviewHandler is a default handler to serve up
// a review page.
func ReviewHandler(c buffalo.Context) error {
	c.Set("title", "ReviewsPage")
	// Get the user's past ratings
	// c.Set("pastRatings", ReviewGetUserRatings(c))
	c.Set("pastUsers", ReviewGetPastUsers(c))
	c.Set("pastReviews", ReviewGetPastReviews(c))
	return c.Render(200, r.HTML("reviewspage.html"))
}

func ReviewGetUserRatings(c buffalo.Context) [][]string {

	ratings := [][]string{
		{"Adele","5", "Pretty Good", "10/23/16"},
		{"Bobbin","4", "Alright", "8/13/12"},
		{"Dash","2", "So-so", "7/23/18"},
		{"Justin","5", "Awesome", "8/23/16"},
		{"Sophie","1", "Not so good", "11/23/16"},
	}
	return ratings
}

func ReviewGetPastUsers(c buffalo.Context) []*models.User {
	curr_user := c.Session().Get("user").(*models.User)
	tx := c.Value("tx").(*pop.Connection)
	sent, err := models.GetRequestsSent(curr_user, tx)
	if err != nil {
		return nil
	}

	all_reqs := []models.Request{}

	for index, req := range *sent {
		if req.Status == 1 {
			all_reqs = append(all_reqs, (*sent)[index])
		}
	}

	users := []*models.User{}
	for _, req := range all_reqs {
		user, err := models.GetUserBySysId(tx, req.ReceiverID)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}
	return users
	//now we have a list of all the request.



	// user_info, err := models.GetInfoByGID(tx, curr_user.GoogleID)
	// if err != nil {
	// 	return nil
	// }
	// isTutor := curr_user.IsTutor
	// list := []string{}
	// if isTutor == 2 {
	// 	list = user_info.GetTutees()
	// } else {
	// 	list = user_info.GetTutors()
	// }

	// users := []*models.User{}
	// for _, Id := range list {

	// 	// check if the useres are tutors and tutees here.
	// 	u, err := models.GetUserByGID(tx, Id)
	// 	if err != nil {
	// 		return nil
	// 	}

	// 	users = append(users, u)
	// }
	// return users
}

func ReviewGetPastReviews(c buffalo.Context) []models.Review {
	curr_user := c.Session().Get("user").(*models.User)
	tx := c.Value("tx").(*pop.Connection)
	reviews, err := models.GetPastReviews(tx, curr_user.GoogleID)
	if err != nil {
		return nil
	}
	return reviews
}
