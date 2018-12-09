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
    languages := ProfileGetLanguages()
	// all the subjects and courses affliated with user
	subjsAndClasses := ProfileGetSubjsAndClasses()
	// slice of strings of all the subjects (Gets the names of all the subjects)
	subjects := ProfileGetSubjects()

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
	c.Set("zip", uinfo.Address)
	// phone number
	c.Set("phone", user.PhoneNumber)
	// email
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

func UpdateProfilePOSTHandler(c buffalo.Context) error {
    fmt.Println("1,2 Buckle my shoe")
    var test map[string][]string
    test = c.Request().Form
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
    //uinfo.SetSubjects([]string{"1","2"})
    uinfo.SetSubjects(GetSubjectsFromCourses(test["Courses[]"]))

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


