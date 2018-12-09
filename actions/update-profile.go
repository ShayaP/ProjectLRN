package actions

import (
    "fmt"
    //"encoding/json"
    //"io/ioutil"

    "github.com/cileonard/lrn/models"


	"github.com/gobuffalo/validate"
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo"
)
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
            //Subjects:   "",
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
	//states := []string{"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", 
	//			  "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", 
	//			  "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", 
	//			  "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", 
	//			  "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}
    */
    languages := ProfileGetLanguages()
    // address of the user 
	//address := ProfileGetAddress(c)
	// get list of all the subjects asociated with the user
	//userSubjects := ProfileGetSubjects()
	// all the subjects and courses affliated with user
	subjsAndClasses := ProfileGetSubjsAndClasses()
	// slice of strings of all the subjects (Gets the names of all the subjects)
	subjects := ProfileGetSubjects()
    /**
    subjects := make([]string, 0, 15)
	for key, _ := range subjsAndClasses {
		subjects = append(subjects, key)
	}*/

	// maps the user courses to all courses
	mapUCtoC := make(map[string]string)
	for _, val := range subjsAndClasses {
		for i:= 0; i < len(val); i++ {
			mapUCtoC[val[i]] = "";
		}
	}

    userCourses := uinfo.GetCourses()
	// Algorithm to put a check on each checkbox course user knows:
	// 	for each subject in subjsAndClasses affliated with user,
	//	Go through the user courses one at a time and	
	// 	run through all the courses under that subject until the user
	//	course is found
    for _, course := range userCourses {
        mapUCtoC[course] = "checked"
    }


    /**
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
	}*/

	// retrieve the languages known by the user
	userLangs := uinfo.GetLanguages()
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
	//c.Set("username", ProfileGetUsername(c))
    // address of user
	//c.Set("street", address[0])
	//c.Set("city", address[1])
	//c.Set("state", address[2])
	//c.Set("states", states)
	c.Set("zip", UpdateProfileGetSelectedZip(c,uinfo.Address))
	//c.Set("zip", address[3])
	// get which zip of the three are the user's zip

	// phone number
	c.Set("phone", user.PhoneNumber)
	// email
	//c.Set("contactEmail", ProfileGetContactEmail(c))
	c.Set("accountEmail", user.Email)
	//Subjects and Languages - help description
	c.Set("subjectDescription", ProfileGetSubjectTip(isTutor))
	c.Set("langDescription", ProfileGetLanguageTip(isTutor))
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

//Description:	returns slice of 3 strings which contains
//				1 string which is selected and the other blanks.
//				In the result slice, 
//					UCSD zip is at index 0
//					SDSU zip is at index 1
//					USD zip is at index 2
func UpdateProfileGetSelectedZip(c buffalo.Context, zip string) [3]string {
	result := [3]string{}
	zipUcsd := "92093"
	zipSdsu := "92182"
	zipUsd := "92110"
	if zip == zipUcsd {
		result[0] = "selected"
	} else {
		result[0] = ""
	}
	if zip == zipSdsu{
		result[1] = "selected"
	} else {
		result[1] = ""
	}
	if zip == zipUsd {
		result[2] = "selected"
	} else {
		result[2] = ""
	}
	return result
}

func UpdateProfilePOSTHandler(c buffalo.Context) error {
    fmt.Println("1,2 Buckle my shoe")
    //test := &models.UpdateProfileForm{}
    //var test models.UpdateProfileForm
    //return c.Render(200, r.JSON(c.Request().Form))
    var test map[string][]string
    test = c.Request().Form
    //fmt.Println(c.Request().Body)
    //body, err := ioutil.ReadAll(c.Request().Body)
    //if err != nil {
    //  return c.Render(500, r.String(err.Error()))
    //}
    //fmt.Println(body)
    //err = json.Unmarshal(body, test)
    //err := c.Bind(test)
    //decoder := json.NewDecoder(c.Request().Body)
    //err = decoder.Decode(&test)
    /**err := nil
    fmt.Println("3,4 Shut the door")
    if err != nil{
        fmt.Println("5,6 This is a dick")
        return c.Render(500, r.String(err.Error()))
    }*/

    fmt.Println("We here boys")
    user := c.Session().Get("user").(*models.User)
	tx := c.Value("tx").(*pop.Connection)
	uinfo, err := models.GetInfoByGID(tx, user.GoogleID)
    newUser := false
    if err != nil{
        //user has not entered any info yet, so uinfo is null
        uinfo = &models.Userinfo{
            Languages:  "",
            Courses:    "",
            Subjects:   "",
            Address:    "",
            GoogleID:   user.GoogleID,
        }
        newUser = true
    }

    user.PhoneNumber = test["PhoneNumber"][0]
    uinfo.Address = test["Location"][0]
    uinfo.SetCourses(test["Courses[]"])
    uinfo.SetLanguages(test["Languages[]"])
    uinfo.SetSubjects([]string{"1","2"})

    //fmt.Println(test["Courses[]"])
    //fmt.Println(test["Languages[]"])

    //return c.Redirect(302,"/update-profile")
    /**
    user.PhoneNumber = test.PhoneNumber
    uinfo.Address = test.Location
    uinfo.SetCourses(test.Courses)
    uinfo.SetLanguages(test.Languages)
    */
    var verrs *validate.Errors
    if newUser{
        verrs, err = uinfo.CreateEntry(tx)
    }else{
        verrs, err = uinfo.UpdateEntry(tx)
    }
	if err != nil { //Error creating the account
		return c.Render(500, r.String(err.Error()))
		//return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Session().Set("errors", verrs)
        fmt.Println(verrs)
        return c.Redirect(302,"/update-profile")
		//return c.Render(200, r.HTML("register.html"))
	}
    verrs, err = user.UpdateEntry(tx)

	if err != nil { //Error creating the account
		return c.Render(500, r.String(err.Error()))
		//return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Session().Set("errors", verrs)
        printverrs(verrs)
        return c.Redirect(302,"/update-profile")
		//return c.Render(200, r.HTML("register.html"))
	}

    fmt.Println("Well fuck")
    //return c.Render(200, r.JSON(c.Request().Form))
    return c.Redirect(302, "/profile")
}

func printverrs(verrs *validate.Errors){
    for _, val := range verrs.Errors{
        fmt.Println(val)
    }
}


