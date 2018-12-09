package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/cileonard/lrn/models"
	"github.com/gobuffalo/pop"
	"fmt"
	"strings"
)

type Name struct {
	First string
}


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

func ReviewPOSTHandler(c buffalo.Context) error {
	rev := &models.Review{}
	n := &Name{}
	tx := c.Value("tx").(*pop.Connection)
	if err := c.Bind(n); err != nil {
		fmt.Println("FIRST ERROR@!!#!#")
		return c.Render(500, r.String(err.Error()))
	}
	first := strings.Split(n.First, " ")[0]
	u, err := models.GetUserByName(tx, first)
	if err != nil {
		return c.Error(401, err)
	}

	if err := c.Bind(rev); err != nil {
		fmt.Println("SECOND ERROR@!!#!#")
		return c.Render(500, r.String(err.Error()))
	}
	curr_user := c.Session().Get("user").(*models.User)
	rev.Reviewer = curr_user.GoogleID
	rev.Reviewee = u.GoogleID
	rev.Astutor = 1

	//push the rev in the db
	verrs, err := tx.ValidateAndCreate(rev)
	if err != nil {
		return c.Error(401, err)
	}

	if verrs.HasAny() {
		fmt.Println("VALIDATION ERRORS!!")
		return c.Redirect(302, "/")
	}

	return c.Redirect(302, "/reviewspage")
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
