package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/cileonard/lrn/models"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"

    "strconv"
    "fmt"
	//"strings"
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

func ReviewPOSTHandler(c buffalo.Context) error {
    //return c.Render(200, r.JSON(c.Request().Form))
	//n := &Name{}
    ratingStr := c.Request().Form["rating"][0]
    review := c.Request().Form["review"][0]
    reviewUsrIDstr := c.Request().Form["reviewedID"][0]

    tx := c.Value("tx").(*pop.Connection)
	//if err := c.Bind(n); err != nil {
	//	fmt.Println("FIRST ERROR@!!#!#")
	//	return c.Render(500, r.String(err.Error()))
	//}
	curr_user := c.Session().Get("user").(*models.User)

    //FIrst we are going to update the new users rating
    reviewedUserID, err := uuid.FromString(reviewUsrIDstr)
    reviewedUser := &models.User{}
    err = tx.Find(reviewedUser, reviewedUserID)
    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }

    rating, err := strconv.Atoi(ratingStr)
    if err != nil{
        return c.Render(500, r.String(err.Error()))
    }
    newrating := (reviewedUser.AvgRating * float32(reviewedUser.NumRatings)) + float32(rating)
    reviewedUser.NumRatings = reviewedUser.NumRatings + 1
    newrating = newrating/float32(reviewedUser.NumRatings)
    reviewedUser.AvgRating = newrating


    //Second we create the review
    rev := &models.Review{
        Rating:         rating,
        Description:    review,
        Reviewer:       curr_user.GoogleID,
        Reviewee:       reviewedUser.GoogleID,
        Astutor:        1,
    }
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
	received, err := models.GetRequestsReceived(curr_user, tx)
	if err != nil {
		return nil
	}

	all_sent_reqs := []models.Request{}
	all_rec_reqs := []models.Request{}

	for index, req := range *sent {
		if req.Status == 1 {
			all_sent_reqs = append(all_sent_reqs, (*sent)[index])
		}
	}
	for index, req := range *received {
		if req.Status == 1 {
			all_rec_reqs = append(all_rec_reqs, (*received)[index])
		}
	}

	users := []*models.User{}
	for _, req := range all_sent_reqs {
		user, err := models.GetUserBySysId(tx, req.ReceiverID)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}
	for _, req := range all_rec_reqs {
		user, err := models.GetUserBySysId(tx, req.SenderID)
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
