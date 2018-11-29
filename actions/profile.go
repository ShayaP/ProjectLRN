package actions

import "github.com/gobuffalo/buffalo"

// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
	var isTutor = true
	c.Set("title", "Profile")
	//display user's name
	c.Set("username", "Bobby Tefla")
	// address of user
	c.Set("street", "645 Mt. Yosemite dr")
	c.Set("city", "San Diego")
	c.Set("state", "CA")
	c.Set("zip", "92084")
	// phone number
	c.Set("phone", "951-842-6895")
	// email
	c.Set("contactEmail", "bobby_tefla@yahoo.com")
	c.Set("accountEmail", "btefla@gmail.com")
	//Subjects and Languages - help description
	if(isTutor){
		c.Set("subjectDescription", "Subjects you can help with")
		c.Set("langDescription", "Languages you are comfortable teaching in")
	} else {
		c.Set("subjectDescription", "Subjects you need help in")
		c.Set("langDescription", "Languages you are comfortable learning in")
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
