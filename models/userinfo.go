package models

import (
	"encoding/json"
	"time"
    "strings"

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
