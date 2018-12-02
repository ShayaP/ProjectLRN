package actions

import "github.com/gobuffalo/buffalo"

// UpdateProfileHandler is a default handler to serve up
// a update-profile page.
func UpdateProfileHandler(c buffalo.Context) error {
	// all the user's personal info that needs to be dynamically set
	// indicates whether user is a tutor
	isTutor := true
	// address of the user 
	address := ProfileGetAddress(c)
	c.Set("title", "Update Profile")
	//display user's name
	c.Set("username", ProfileGetUsername(c))
	// address of user
	c.Set("street", address[0])
	c.Set("city", address[1])
	//c.Set("state", address[2])
	c.Set("zip", address[3])
	// phone number
	c.Set("phone", ProfileGetPhone(c))
	// email
	c.Set("contactEmail", ProfileGetContactEmail(c))
	c.Set("accountEmail", ProfileGetAccountEmail(c))
	//Subjects and Languages - help description
	c.Set("subjectDescription", ProfileGetSubjectTip(c, isTutor))
	c.Set("langDescription", ProfileGetLanguageTip(c, isTutor))
	return c.Render(200, r.HTML("update-profile.html"))
}
