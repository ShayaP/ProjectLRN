package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/cileonard/lrn/models"
	"fmt"
	"strconv"
	"strings"
	)

type Search struct {
	Name string
	Languages string
	Location int
	Topics string
}

// BrowseProfilesHandler is a default handler to serve up
// a home page.
func BrowseProfilesHandler(c buffalo.Context) error {
        c.Set("title", "Browse")
        // c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
        // c.Set("results", u)
        return c.Render(200, r.HTML("browseprofiles.html"))
}

func BrowseProfilesPOSTHandler(c buffalo.Context) error {
    c.Set("title", "Browse")
    s := &Search{}
	if err := c.Bind(s); err != nil {
		return c.Render(500, r.String(err.Error()))
	}
    // return c.Render(200, r.JSON(s))
	name := s.Name
	fmt.Println(name)
	tx := c.Value("tx").(*pop.Connection)

	curr_user := c.Session().Get("user").(*models.User)
	isTutor := curr_user.IsTutor
	if name == "" {
		lang := s.Languages
		loc := strconv.Itoa(s.Location)
		topics := s.Topics
		list, err := models.GetFilteredUsers(tx, loc, lang, topics)
		if err != nil {
			return c.Error(401, err)
		}
		fmt.Println()
		fmt.Println(list)
		users := []*models.User{} 
		courses := [][]string{}
		langs := [][]string{}
		for _, userinfo := range list {
			// check if the useres are tutors and tutees here.
			u, err := models.GetUserByGID(tx, userinfo.GoogleID)
			fmt.Println("\n")
			fmt.Println(u.FirstName)
			fmt.Println("\n")
			if err != nil {
				return c.Error(401, err)
			}

			//automatically show only the other group of people 
			if u.IsTutor == isTutor {
				continue
			}

			//remove the current user from the list
			if u.GoogleID == curr_user.GoogleID {
				continue
			}

			users = append(users, u)
			langs = append(langs, userinfo.GetLanguages())
			courses = append(courses, userinfo.GetCourses())
		}
		fmt.Println("\n")
		fmt.Println(users)
		fmt.Println("\n")
		c.Set("topic", topics)
		c.Set("results", users)
		c.Set("courses", courses)
		c.Set("langs", langs)
	} else {
        name = strings.Split(name, " ")[0]
		u, err := models.GetUserByName(tx, name, isTutor)
		if err != nil {
		    users := []*models.User{} 
		    courses := [][]string{}
		    langs := [][]string{}
            topics := s.Topics
            c.Set("topic", topics)
		    c.Set("results", users)
		    c.Set("courses", courses)
		    c.Set("langs", langs)
			return c.Render(200, r.HTML("browseprofiles.html"))
            //return c.Error(401, err)
		}
		fmt.Println("222222")
        temp := []*models.User{u}
		c.Set("topic", "Variety")
		info, err := models.GetInfoByGID(tx, u.GoogleID)
		if err != nil {
			return c.Error(401, err)
		}
		courses := [][]string{}
		langs := [][]string{}
		langs = append(langs, info.GetLanguages())
		courses = append(courses, info.GetCourses())
		c.Set("results", temp)
		c.Set("courses", courses)
		c.Set("langs", langs)
	}
    // c.Set("results", []string{"Samantha B","John S","Gary G","Aaron R","Michael N"})
    return c.Render(200, r.HTML("browseprofiles.html"))
}
