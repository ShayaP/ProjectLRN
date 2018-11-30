package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	PhoneNumber  string    `json:"phone_number" db:"phone_number"`
	GoogleID     string    `json:"google_id" db:"google_id"`
	Email        string    `json:"email" db:"email"`
	Gender       int       `json:"gender" db:"gender"`
	OtherSpecify string    `json:"other_specify" db:"other_specify"`
    AvgRating    float32   `json:"avg_rating" db:"avg_rating"`
    NumRatings   int       `json:"num_ratings" db:"num_ratings"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: u.PhoneNumber, Name: "PhoneNumber"},
		&validators.StringIsPresent{Field: u.GoogleID, Name: "GoogleID"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.IntIsPresent{Field: u.Gender, Name: "Gender"},
		&validators.StringIsPresent{Field: u.OtherSpecify, Name: "OtherSpecify"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.FirstName = strings.Title(u.FirstName)
	u.LastName = strings.Title(u.LastName)
	return tx.ValidateAndCreate(u)
}

//func GetUser(tx) (*User, error){

//}
