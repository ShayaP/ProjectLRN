package models

import (
	"encoding/json"
	// "time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Review struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Rating      int       `json:"rating" db:"rating"`
	Description string    `json:"description" db:"description"`
	Reviewer    string 	  `json:"reviwer" db:"reviewer"`
	Reviewee    string 	  `json:"reviewee" db:"reviewee"`
	Astutor     int       `json:"astutor" db:"astutor"`
}

// String is not required by pop and may be deleted
func (r Review) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Reviews is not required by pop and may be deleted
type Reviews []Review

// String is not required by pop and may be deleted
func (r Reviews) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Review) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: r.Rating, Name: "Rating"},
		&validators.StringIsPresent{Field: r.Description, Name: "Description"},
		&validators.IntIsPresent{Field: r.Astutor, Name: "Astutor"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Review) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Review) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func GetPastReviews(tx *pop.Connection, gid string) ([]Review, error) {
	query := tx.RawQuery("SELECT * FROM reviews WHERE reviewer = ?", gid)
	reviews := &[]Review{}
	if err := query.All(reviews); err != nil {
		return nil, err
	}
	return *reviews, nil

}
