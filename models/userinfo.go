package models

import (
	"encoding/json"
	"time"
    "strings"
    "fmt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

var tokenSplit = "/"

type Userinfo struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Languages string    `json:"languages" db:"languages"`
	Subjects  string    `json:"subjects" db:"subjects"`
	Courses   string    `json:"courses" db:"courses"`
	Address   string    `json:"address" db:"address"`
    GoogleID  string    `json:"google_id" db:"google_id"`
    Tutors	  string 	`json:"tutors" db:"tutors"`
    Tutees	  string 	`json:"tutees" db:"tutees"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (u Userinfo) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Userinfoes is not required by pop and may be deleted
type Userinfoes []Userinfo

// String is not required by pop and may be deleted
func (u Userinfoes) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *Userinfo) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Languages, Name: "Languages"},
		&validators.StringIsPresent{Field: u.Subjects, Name: "Subjects"},
		&validators.StringIsPresent{Field: u.Courses, Name: "Courses"},
		&validators.StringIsPresent{Field: u.Address, Name: "Address"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *Userinfo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *Userinfo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *Userinfo) UpdateEntry(tx *pop.Connection) (*validate.Errors, error){
    return tx.ValidateAndUpdate(u)
}


func (u *Userinfo) CreateEntry(tx *pop.Connection) (*validate.Errors, error){
    return tx.ValidateAndCreate(u)
}

func (u *Userinfo) GetLanguages() []string {
    langs := decodeStringToIds(u.Languages)
    return langs
}

func (u *Userinfo) SetLanguages(langs []string) {
    u.Languages = encodeIdsToString(langs)
}

func (u *Userinfo) GetCourses() []string {
    courses := decodeStringToIds(u.Courses)
    return courses
}

func (u *Userinfo) SetCourses(courses []string) {
    u.Courses = encodeIdsToString(courses)
}

func (u *Userinfo) GetSubjects() []string {
    courses := decodeStringToIds(u.Subjects)
    return courses
}

func (u *Userinfo) SetSubjects(courses []string) {
    u.Subjects = encodeIdsToString(courses)
}



func (u *Userinfo) SetTutors(tutors []string) {
    u.Tutors = encodeIdsToString(tutors)
}

func (u *Userinfo) GetTutors() []string {
    tutors := decodeStringToIds(u.Tutors)
    return tutors
}

func (u *Userinfo) SetTutees(tutors []string) {
    u.Tutees = encodeIdsToString(tutors)
}

func (u *Userinfo) GetTutees() []string{
    tutees := decodeStringToIds(u.Tutees)
    return tutees
}

func decodeStringToIds(instring string) []string{
    list := strings.Split(strings.Trim(instring, tokenSplit), tokenSplit)
    return list
}

func encodeIdsToString(arr []string) string{
    var b strings.Builder
    b.WriteString(tokenSplit)
    b.WriteString(strings.Join(arr, tokenSplit))
    b.WriteString(tokenSplit)
    return b.String()
}

func GetInfoByGID(tx *pop.Connection, gid string) (*Userinfo, error){
	query := tx.RawQuery("SELECT * FROM userinfoes WHERE google_id = ?", gid)
	u := &Userinfo{}
	if err := query.First(u); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func GetInfoBySysID(tx *pop.Connection, id uuid.UUID) (*Userinfo, error){
	query := tx.RawQuery("SELECT * FROM userinfoes WHERE user_id = ?", id)
	u := &Userinfo{}
	if err := query.First(u); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func GetFilteredUsers(tx *pop.Connection, location string, language string, subject string) ([]Userinfo, error) {
	query := tx.RawQuery("SELECT * FROM userinfoes")
	userList := &[]Userinfo{}
	if err := query.All(userList); err != nil {
		return nil, err
	}

	returnUsers := []Userinfo{}
	for index, user := range *userList {
		langs := user.GetLanguages()
		loc := user.Address
		subs := user.GetSubjects()

		boolLang := false
		for _, l := range langs {
			fmt.Println("\n")
			fmt.Println(language)
			fmt.Println(l)
			fmt.Println("\n")
			if l == language {
				boolLang = true
				break
			}
		}

		// check location
		boolLoc := (loc == location)

		boolSub := false
		for _, s := range subs {
			fmt.Println("\n")
			fmt.Println(subject)
			fmt.Println(s)
			fmt.Println("\n")
			if s == subject {
				boolSub = true
				break
			}
		}

		if boolLang == true && boolSub == true && boolLoc == true {
			returnUsers = append(returnUsers, (*userList)[index])
		}
	}
	fmt.Println("\n")
	fmt.Println(returnUsers)
	fmt.Println("\n")
	return returnUsers, nil
}
