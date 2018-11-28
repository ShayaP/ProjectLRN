package models

import (
	"errors"
)

func GetUser(tokenID string) (*User, error) {
	u := &User{}
	err := DB.Find(u, tokenID)
	if err != nil {
		return u, err
	} else {
		return u, nil
	}
}

func CreateUserEntry(user *User) error {
	verrs, err := DB.ValidateAndCreate(user)
	if len(verrs.Errors) != 0 {
		return errors.New(string(len(verrs.Errors)) + " erros in validation")
	}
	return err
}
