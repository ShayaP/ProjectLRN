package actions

import "github.com/gobuffalo/buffalo"


// ReviewHandler is a default handler to serve up
// a review page.
func ReviewHandler(c buffalo.Context) error {
	c.Set("title", "ReviewsPage")
	// Get the user's past ratings
	c.Set("pastRatings", ReviewGetUserRatings(c))
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
