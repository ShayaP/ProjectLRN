package models

import (
	"log"
	//"fmt"
	//"encoding/json"
	//"time"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	//"github.com/gobuffalo/uuid"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

/**
type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	PhoneNumber  string
	Email        string
	UserName     string
	UserPassword string
	//Image BLOB
	// Location BLOB
	Gender       int
	OtherSpecify string
	TutorId      int
	TuteeId      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
*/
func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
	//fmt.Println("Hi!");
}
