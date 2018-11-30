package actions

import "github.com/gobuffalo/buffalo"

// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
	// all the user's personal info that needs to be dynamically set
	var isTutor = true
	var username = "Bobby Tefla"
	var street = "645 Mt. Yosemite dr"
	var city = "San Diego"
	var state = "CA"
	var zip = "92084"
	var phone = "951-842-6895"
	var contactEmail = "bobby_tefla@yahoo.com"
	var accountEmail = "btefla@gmail.com"
	var ttrSubjectDescription = "Subjects you can help with"
	var tteSubjectDescription = "Subjects you need help in"
	var ttrLangDescription = "Languages you are comfortable teaching in"
	var tteLangDescription = "Languages you are comfortable learning in"
	c.Set("title", "Profile")
	//display user's name
	c.Set("username", username)
	// address of user
	c.Set("street", street)
	c.Set("city", city)
	c.Set("state", state)
	c.Set("zip", zip)
	// phone number
	c.Set("phone", phone)
	// email
	c.Set("contactEmail", contactEmail)
	c.Set("accountEmail", accountEmail)
	//Subjects and Languages - help description
	if(isTutor){
		c.Set("subjectDescription", ttrSubjectDescription)
		c.Set("langDescription", ttrLangDescription)
	} else {
		c.Set("subjectDescription", tteSubjectDescription)
		c.Set("langDescription", tteLangDescription)
	}
	//Subjects 
	// list of the subjects that a student might need help in 
	var mainSubjects = []string{"Mathematics", "Chemistry"}
	// pairs the subject and specific classes under that subject together
	subjectsAndClasses := make(map[string][]string)
	subjectsAndClasses["Mathematics"] = []string{"Linear Algebra","Differential Equations"}
	subjectsAndClasses["Chemistry"] = []string{"Introduction to Chemistry"}
	c.Set("mainSubjects", mainSubjects)
	c.Set("subjectsAndClasses", subjectsAndClasses)
	//Languages 
	// these variables split the languages into two equal size parts (for styling)
	var langsPart1 = []string{"English"};
	var langsPart2 = []string{"Chinese"};
	c.Set("languages1", langsPart1)
	c.Set("languages2", langsPart2)
	return c.Render(200, r.HTML("profile.html"))
}
