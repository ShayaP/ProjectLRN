package models

import (
	"encoding/json"
	"time"
    "strconv"
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


func (u *Userinfo) GetLanguages() []int {
    list := strings.Split(u.Languages, tokenSplit)
    var languages = []int{}

    for _,i := range list {
        j, _ := strconv.Atoi(i)

        languages = append(languages, j)
    }
    return languages

}

