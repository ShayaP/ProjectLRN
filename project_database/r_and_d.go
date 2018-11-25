package main

import (
	"fmt"

	"github.com/cileonard/lrn/models"
)

func main() {
	jessica := models.User{Title: "Ms.", FirstName: "Jessica", LastName: "Jones", Bio: "Private security, super hero."}
	fmt.Println("%s\n", jessica.FirstName)

	//_, err = *DB.ValidateAndSave(&jessica)

	//if err != nil {
	//	log.Panic(err)
	//}
}
