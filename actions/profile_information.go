package actions

import (
    //"fmt"
    "sort"

    "github.com/cileonard/lrn/models"

    //"github.com/gobuffalo/buffalo"
)

var subjs_and_classes map[string][]string
var all_langs []string

func init(){

    outputMap := make(map[string][]string, len(models.StaticCourses))
    for subj, courses := range models.StaticCourses{
        outputMap[subj] = make([]string, len(courses))
        for i, course := range courses {
            outputMap[subj][i] = course.Course
        }
        sort.Strings(outputMap[subj])
    }
    subjs_and_classes = outputMap

	all_langs = make([]string, len(models.StaticLangs))
    cnt := 0
    for _,l := range models.StaticLangs {
        all_langs[cnt] = l
        cnt++
    }
    sort.Strings(all_langs)
    //fmt.Println(all_langs)
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
func ProfileGetSubjectTip(isTutor bool) string {
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
func ProfileGetLanguageTip(isTutor bool) string {
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
func ProfileGetSubjsAndClasses() map[string][]string {
	// pairs the user's subject and specific classes under that subject together
	return subjs_and_classes
}

//Name: 	ProfileGetSubjects
//Purpose: 	gets the subjects associated with the user (Ex: if user needs
//		tutoring in Linear Algebra, the subject would be Mathematics)
//Description:	returns a pre-fabricated list of subjects
//Parameters:	c - a context
//Return:	returns a pre-fabricated list of subjects
func ProfileGetSubjects() []string {
	subjectKeys := make([]string, 0, len(subjs_and_classes))
	for key, _ := range subjs_and_classes {
		subjectKeys = append(subjectKeys, key)
	}
    sort.Strings(subjectKeys)
	return subjectKeys
}

//Name: 	ProfileGetLanguagess
//Purpose: 	gets the languages associated with the user
//Description:	returns a pre-fabricated list of languages
//Parameters:	c - a context
//Return:	returns a pre-fabricated list of languages
func ProfileGetLanguages() []string {
	return all_langs
}

