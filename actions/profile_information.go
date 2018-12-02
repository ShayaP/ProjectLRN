package actions

import "github.com/gobuffalo/buffalo"

//Name: 	ProfileGetUserame
//Purpose: 	gets the user's name to display
//Description:	returns a pre-fabricated user name
//Parameters:	c - a context
//Return:	returns the user's name as a string
func ProfileGetUsername(c buffalo.Context) string {
	// all the user's personal info that needs to be dynamically set
	var username = "Bobby Tefla"
	return username
}

//Name: 	ProfileGetAddress
//Purpose: 	gets the user's address (street, city, state, zip) to display
//Description:	returns a pre-fabricated address
//Parameters:	c - a context
//Return:	returna string array of the users address in order:
//		["street","city", "state", "zip"]
func ProfileGetAddress(c buffalo.Context) []string {
	// the user's address
	street := "645 Mt. Yosemite dr"
	city := "San Diego"
	state := "CA"
	zip := "92084"
	address := []string{street, city, state, zip}
	return address
}

//Name: 	ProfileGetPhone
//Purpose: 	gets the user's phone number to
//Description:	returns a pre-fabricated phone number
//Parameters:	c - a context
//Return:	returns phone number as string in format: xxx-xxx-xxxx
func ProfileGetPhone(c buffalo.Context) string {
	// user's phone number
	phone := "951-842-6895"
	return phone
}

//Name: 	ProfileGetContactEmail
//Purpose: 	gets the user's contact email to display
//Description:	returns a pre-fabricated contact email
//Parameters:	c - a context
//Return:	returns contact email as a string
func ProfileGetContactEmail(c buffalo.Context) string {
	// user's contact email
	contactEmail := "bobby_tefla@yahoo.com"
	//display user's contact email
	return contactEmail
}

//Name: 	ProfileGetAccountEmail
//Purpose: 	gets the user's account email to display
//Description:	returns a pre-fabricated account email
//Parameters:	c - a context
//Return:	returns account email as a string
func ProfileGetAccountEmail(c buffalo.Context) string {
	// user's account email
	accountEmail := "btefla@gmail.com"
	return accountEmail
}

//Name: 	ProfileGetSubjectTip
//Purpose: 	gets the description of what the subjects section in
//		profile means to the user (changes if tutor or tutee)
//Description:	returns a description relevant to tutors if isTutor is true.
//		Otherwise, returns a description relevant to tutees
//Parameters:	c - a context
//		isTutor - boolean that is true if the user is a tutor and
//			false otherwise
//Return:	returns a description relevant to tutors if the user is
//		a tutor or a description relevant to tutees if the user
//		is a tutee
func ProfileGetSubjectTip(c buffalo.Context, isTutor bool) string {
	// description of what the subjects section in profile means to tutors
	ttrSubjectDescription := "Subjects you can help with"
	// description of what the subjects section in profile means to tutees
	tteSubjectDescription := "Subjects you need help in"

	if isTutor {
		return ttrSubjectDescription
	} else {
		return tteSubjectDescription
	}
}

//Name: 	ProfileGetLangaugeTip
//Purpose: 	gets the description of what the langauges section in
//		profile means to the user (changes if tutor or tutee)
//Description:	returns a description relevant to tutors if isTutor is true.
//		Otherwise, returns a description relevant to tutees
//Parameters:	c - a context
//		isTutor - boolean that is true if the user is a tutor and
//			false otherwise
//Return:	returns a description relevant to tutors if the user is
//		a tutor or a description relevant to tutees if the user
//		is a tutee
func ProfileGetLanguageTip(c buffalo.Context, isTutor bool) string {
	// description of what the languages section in profile means to tutors
	var ttrLangDescription = "Languages you are comfortable teaching in"
	// description of what the languages section in profile means to tutees
	var tteLangDescription = "Languages you are comfortable learning in"

	if isTutor {
		return ttrLangDescription
	} else {
		return tteLangDescription
	}
}

//Name: 	ProfileGetSubjsAndClasses
//Purpose: 	gets the subjects and classes associated with the user
//Description:	returns a pre-fabricated map of key subjects and list of classes
//			under each subejct
//Parameters:	c - a context
//Return:	returns a pre-fabricated map of key subjects and list of classes
//			under each subject
func ProfileGetSubjsAndClasses(c buffalo.Context) map[string][]string {
	// pairs the user's subject and specific classes under that subject together
	subjectsAndClasses := make(map[string][]string)
	subjectsAndClasses["Mathematics"] = []string{"Linear Algebra", "Differential Equations"}
	subjectsAndClasses["Chemistry"] = []string{"Introduction to Chemistry"}

	return subjectsAndClasses
}

//Name: 	ProfileGetSubjects
//Purpose: 	gets the subjects associated with the user (Ex: if user needs
//		tutoring in Linear Algebra, the subject would be Mathematics)
//Description:	returns a pre-fabricated list of subjects
//Parameters:	c - a context
//Return:	returns a pre-fabricated list of subjects
func ProfileGetSubjects(c buffalo.Context) []string {
	//  delegate to ProfileGetSubjectAndClasses to get a map of
	// subjects and classes  then extract the keys and return them
	//subjsAndClasses := make(map[string][]string)
	subjsAndClasses := ProfileGetSubjsAndClasses(c)
	// store the keys in the map (Assumed max keys is 20)
	subjectKeys := make([]string, 0, 10)
	for key, _ := range subjsAndClasses {
		subjectKeys = append(subjectKeys, key)
	}

	return subjectKeys
}

//Name: 	ProfileGetLanguagess
//Purpose: 	gets the languages associated with the user
//Description:	returns a pre-fabricated list of languages
//Parameters:	c - a context
//Return:	returns a pre-fabricated list of languages
func ProfileGetLanguages(c buffalo.Context) []string {
	languages := []string{"English", "Chinese", "Japanese", "Korean"}
	return languages
}
