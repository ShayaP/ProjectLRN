package actions

import "github.com/gobuffalo/buffalo"

// UpdateProfileHandler is a default handler to serve up
// a update-profile page.
func UpdateProfileHandler(c buffalo.Context) error {
    user := c.Session().Get("user").(*models.User)
	tx := c.Value("tx").(*pop.Connection)
	uinfo, err := models.GetInfoByGID(tx, user.GoogleID)
    if err != nil{
        //user has not entered any info yet, so uinfo is null
        uinfo = &models.Userinfo{
            Languages:  "",
            Courses:    "",
            Subjects:   "",
            Address:    "",
        }
    }
	// all the user's personal info that needs to be dynamically set
	// indicates whether user is a tutor
	isTutor := (user.IsTutor == 2)
	// comprehensive list of languages user can choose
	/**
    languages := []string{"English", "Arabic", "Armenian", "Austronesian", "Chinese", "French",
				"German", "Hindi", "Japanese", "Korean", "Persian", "Portugese",
				"Punjabi", "Russian","Spanish", "Tagalog", "Tai-Ka", "Vietnamese"}
	// comprehensive list of courses user can choose
	var subjsAndClasses = map[string][]string{
		"Applied Arts" : {"Digital Media", "Film Production","Photography"},
		"Biology": {"Botany", "Marine Biology", "Zoology"},
		"Chemistry": {"Introduction to Chemistry", "Organic Chemistry"},
	}
	//comprehensive list of states user can choose from
	states := []string{"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", 
				  "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", 
				  "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", 
				  "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", 
				  "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}
    */
	// address of the user 
	//address := ProfileGetAddress(c)
	// get list of all the subjects asociated with the user
	userSubjects := ProfileGetSubjects(c)
	// all the subjects and courses affliated with user
	userSubsAndClasses := ProfileGetSubjsAndClasses(c)
	// slice of strings of all the subjects (Gets the names of all the subjects)
	subjects := make([]string, 0, 15)
	for key, _ := range subjsAndClasses {
		subjects = append(subjects, key)
	}

	// maps the user courses to all courses
	mapUCtoC := make(map[string]string)
	for _, val := range subjsAndClasses {
		for i:= 0; i < len(val); i++ {
			mapUCtoC[val[i]] = "";
		}
	}

	// Algorithm to put a check on each checkbox course user knows:
	// 	for each subject in subjsAndClasses affliated with user,
	//	Go through the user courses one at a time and	
	// 	run through all the courses under that subject until the user
	//	course is found
	for i:=0; i < len(userSubjects); i++ {
		// all courses under the subject affliated with user
		var courses = subjsAndClasses[userSubjects[i]]
		// all the courses the user took
		var userCourses = userSubsAndClasses[userSubjects[i]]
		for j:= 0; j < len(userCourses); j++ {
			//search over all the course until find a match
			for k := 0; k < len(courses); k++ {
				if userCourses[j] == courses[k] {
					mapUCtoC[userCourses[j]] = "checked"
					break;
				}
			}
		}
	}

	// retrieve the languages known by the user
	userLangs := ProfileGetLanguages(c)
	// Algorithm to put a check on each checkbox language user knows:
	// 	for each language in languages, create a key, value pair k,v
	// 	such that k = language and v = "checked" if the language is
	// 	in userLangs and v = "" otherwise
	// Step 1) initialize the mapping of user known languages to the
	// complete set of languages so keys are all the languages and v is
	// defaulted to ""
	mapULtoL := make(map[string]string)
	for i := 0; i < len(languages); i++ {
		mapULtoL[languages[i]] = ""
	}

	// Step 2) access values in mapULtoL using values in userLangs and set
	//	to "checked"
	for i := 0; i < len(userLangs); i++ {
		mapULtoL[userLangs[i]] = "checked"
	}

	c.Set("title", "Update Profile")
	//display user's name
	c.Set("username", ProfileGetUsername(c))
	// address of user
	c.Set("street", address[0])
	c.Set("city", address[1])
	c.Set("state", address[2])
	c.Set("states", states)
	c.Set("zip", address[3])
	// phone number
	c.Set("phone", ProfileGetPhone(c))
	// email
	c.Set("contactEmail", ProfileGetContactEmail(c))
	c.Set("accountEmail", ProfileGetAccountEmail(c))
	//Subjects and Languages - help description
	c.Set("subjectDescription", ProfileGetSubjectTip(c, isTutor))
	c.Set("langDescription", ProfileGetLanguageTip(c, isTutor))
	//Set subjects options
	c.Set("subjects", subjects)
	c.Set("subjsAndClasses", subjsAndClasses)
	c.Set("courseChecked", mapUCtoC)
	//Set the languages options (split in half for styling)
	c.Set("languages1", languages[:len(languages)/2])
	c.Set("languages2", languages[len(languages)/2:len(languages)])
	c.Set("checked", mapULtoL)
	return c.Render(200, r.HTML("update-profile.html"))
}

func UpdateProfileHandlerPOSTHandler(c buffalo.Context) error {
    return nil
}




