package main

import (
	"fmt"

	"github.com/cileonard/lrn/crud"
)

func main() {
	var validUser *crud.User
	var invalidUser *crud.User
	var findUser int

	validUser = crud.CreateUser("Josh", "Will", "123-456-1234", "myemail@uscd.edu", "jweezy", "123", 0, "")

	if validUser == nil {
		fmt.Println("TEST FAILED TO CREATE VALID USER")
		return
	}

	fmt.Println("VALID USER INSERSTION PASSED")

	invalidUser = crud.CreateUser("", "Will", "123-456-1234", "myemail@uscd.edu", "jweezy", "123", 0, "")

	if invalidUser != nil {
		fmt.Println("TEST FAILED TO REJECT INVALID USER")
		return
	}

	fmt.Println("INVALID USER INSERSTION PASSED")

	findUser = crud.GetUser(validUser.ID)

	if findUser == 0 {
		fmt.Println("TEST FAILED TO FIND VALID USER")
		return
	}

	fmt.Println("VALID USER SEARCH PASSED")

}
