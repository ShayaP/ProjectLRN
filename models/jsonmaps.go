package models
import(
    "encoding/json"
    //"fmt"
	"github.com/gobuffalo/packr"
)

//Course file is a map[string][]Course
//key string is a subject, Each course has its own stuff
type CourseJSON struct {
    ID      string  `json:"id"`
    Course  string  `json:"course"`
}
var StaticCourses   map[string][]CourseJSON
var StaticLangs     map[string]string

func init() {
    box := packr.NewBox("../static")

    langBytes, _ := box.Find("language.json")
    //var langMap Langarray
    json.Unmarshal(langBytes, &StaticLangs)

    courseBytes, _ := box.Find("course.json")
    //var courseMap Coursearray
    json.Unmarshal(courseBytes, &StaticCourses)
    //StaticCourses = courseMap.Courses
    //StaticLangs = langMap.Languages
}
