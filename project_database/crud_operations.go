// Data Querying Class for CSE 110
package crud

import (
	"fmt"

	"github.com/cileonard/lrn/models"
	"github.com/gofrs/uuid"
)

func createUser(firstName string, lastName string, phoneNumber string, email string,
	userName string, userPassword string, gender int, otherSpecify string) *User {

	// Check that fistName, lastName, email, and gender are not NULL
	if len(firstName) == 0 || len(lastName) == nil || len(email) == nil || (gender != 1 && gender != 0 && gender != -1) {
		fmt.ErrorF("NOT NULL value is not implimeneted!")
		return nil
	}

	// Create User struct
	newUser := models.User{FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber,
		Email: email, UserName: userName, UserPassword: userPassword, Gender: gender, OtherSpecify: otherSpecify}

	// Insert new User into the database
	_, err = DB.ValidateAndSave(&newUser)

	// Check if insertion was successful
	if err != nil {
		fmt.ErrorF("INSERTION FAILURE!!\n")
		fmt.Printf("%v\n", err)
		return nil
	}

	return *newUser
}

func getUser(id uuid.UUID) *User {

	searchUser := models.User{}
	err = DB.Find(&searchUser, id)

	if err != nil {
		fmt.ErrorF("ERROR!\n")
		fmt.Printf("%v\n", err)
		return nil

	} else {
		fmt.Print("Success!\n")
		return searchUser
	}
}
