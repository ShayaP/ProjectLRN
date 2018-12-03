package actions

import (
    "strings"
    "fmt"
    "github.com/cileonard/lrn/models"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop"
)
// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
    user := c.Session().Get("user").(*models.User)
	tx := c.Value("tx").(*pop.Connection)
	uinfo, err := models.GetInfoByGID(tx, user.GoogleID)
    if err != nil {
        //user has not created any information yet, so update profile
        fmt.Println(err.Error())
        return c.Redirect(302, "/update-profile")
    }
    // all the user's personal info that needs to be dynamically set
	// indicates whether user is a tutor
	isTutor := (user.IsTutor == 2)
	// address of the user
	//address := ProfileGetAddress(c)
	c.Set("title", "Profile")
	//display user's name
    var b strings.Builder
    b.WriteString(user.FirstName)
    b.WriteString(" ")
    b.WriteString(user.LastName)
	c.Set("username", b.String())
    b.Reset()
	// address of user
	//c.Set("street", address[0])
	//c.Set("city", address[1])
	//c.Set("state", address[2])
	//c.Set("zip", address[3])
	c.Set("zip", uinfo.Address)
	// phone number
	c.Set("phone", user.PhoneNumber)
	// email
	//c.Set("contactEmail", ProfileGetContactEmail(c))
	c.Set("accountEmail", user.Email)
	//Subjects and Languages - help description
	c.Set("subjectDescription", ProfileGetSubjectTip(isTutor))
	c.Set("langDescription", ProfileGetLanguageTip(isTutor))
	//Subjects
	//c.Set("mainSubjects", ProfileGetSubjects())
	//c.Set("subjectsAndClasses", ProfileGetSubjsAndClasses())
	//c.Set("mainSubjects", uinfo.GetSubjects())
    c.Set("subjectsAndClasses", uinfo.GetCourses())
    
    //Languages
	// these variables split the languages into two equal size parts (for styling)
	languages := uinfo.GetLanguages()
	// styling: used to ensure that there are more or equal number of
	// languages displayed in the first of the two containers in profile
	midpoint := 0
	// when the number of languages known is 2k+1, put k+1 of them in the
	// first half
	if len(languages)%2 != 0 {
		midpoint = len(languages)/2 + 1
	} else {
		// when the number of languages known is 2k, put k of them in the
		// first half
		midpoint = len(languages) / 2
	}
	c.Set("languages1", languages[:midpoint])
	c.Set("languages2", languages[midpoint:len(languages)])
	return c.Render(200, r.HTML("profile.html"))
}
