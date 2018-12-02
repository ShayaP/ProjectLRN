package actions

import "github.com/gobuffalo/buffalo"

// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
	// all the user's personal info that needs to be dynamically set
	// indicates whether user is a tutor
	isTutor := true
	// address of the user 
	address := ProfileGetAddress(c)
	c.Set("title", "Profile")
	//display user's name
	c.Set("username", ProfileGetUsername(c))
	// address of user
	c.Set("street", address[0])
	c.Set("city", address[1])
	c.Set("state", address[2])
	c.Set("zip", address[3])
	// phone number
	c.Set("phone", ProfileGetPhone(c))
	// email
	c.Set("contactEmail", ProfileGetContactEmail(c))
	c.Set("accountEmail", ProfileGetAccountEmail(c))
	//Subjects and Languages - help description
	c.Set("subjectDescription", ProfileGetSubjectTip(c, isTutor))
	c.Set("langDescription", ProfileGetLanguageTip(c, isTutor))
	//Subjects
	c.Set("mainSubjects", ProfileGetSubjects(c))
	c.Set("subjectsAndClasses", ProfileGetSubjsAndClasses(c))
	//Languages
	// these variables split the languages into two equal size parts (for styling)
	languages := ProfileGetLanguages(c)
	// styling: used to ensure that there are more or equal number of 
	// languages displayed in the first of the two containers in profile 
	midpoint := 0
	// when the number of languages known is 2k+1, put k+1 of them in the 
	// first half
	if len(languages) % 2 != 0 {
		midpoint = len(languages)/2 + 1
	} else {
		// when the number of languages known is 2k, put k of them in the 
		// first half
		midpoint = len(languages)/2
	}
	c.Set("languages1", languages[:midpoint])
	c.Set("languages2", languages[midpoint:len(languages)])
	return c.Render(200, r.HTML("profile.html"))
}
