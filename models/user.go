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

/**
 *  IsTutor: 1 if tutee, 2 if tutor
 *
 *
 */
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
	IsTutor      int       `json:"is_tutor" db:"is_tutor"`
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
		&validators.EmailLike{Field: u.Email, Name: "Email"},
		&validators.IntIsPresent{Field: u.Gender, Name: "Gender"},
		&validators.StringIsPresent{Field: u.OtherSpecify, Name: "OtherSpecify"},
		&validators.IntIsPresent{Field: u.IsTutor, Name: "IsTutor"},
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
	u.AvgRating = 0
	u.NumRatings = 0
	return tx.ValidateAndCreate(u)
}

func (u *User) UpdateEntry(tx *pop.Connection) (*validate.Errors, error){
    return tx.ValidateAndUpdate(u)
}

func GetUserByGID(tx *pop.Connection, gid string) (*User, error) {
	query := tx.RawQuery("SELECT * FROM users WHERE google_id = ?", gid)
	u := &User{}
	if err := query.First(u); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func GetUserByName(tx *pop.Connection, name string) (*User, error) {
	query := tx.RawQuery("SELECT * FROM users WHERE first_name = ?", name)
	u := &User{}
	if err := query.First(u); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

// func GetUserByLocation(tx *pop.Connection, gid string) (*User, error) {
// 	query := tx.RawQuery("SELECT * FROM users WHERE google_id = ?", gid)
// 	u := &User{}
// 	if err := query.First(u); err != nil {
// 		return nil, err
// 	} else {
// 		return u, nil
// 	}
// }

// func GetUserByLanguage(tx *pop.Connection, gid string) (*User, error) {
// 	query := tx.RawQuery("SELECT * FROM users WHERE google_id = ?", gid)
// 	u := &User{}
// 	if err := query.First(u); err != nil {
// 		return nil, err
// 	} else {
// 		return u, nil
// 	}
// }

// func GetUserByTopic(tx *pop.Connection, gid string) (*User, error) {
// 	query := tx.RawQuery("SELECT * FROM users WHERE google_id = ?", gid)
// 	u := &User{}
// 	if err := query.First(u); err != nil {
// 		return nil, err
// 	} else {
// 		return u, nil
// 	}
// }
