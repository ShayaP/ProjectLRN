package models

import (
    //"testing"
    "fmt"
    "time"

	"github.com/gobuffalo/uuid"
)

func (ms *ModelSuite) Test_CreateUserEntry(){
    validUser := &User{uuid.Nil, time.Now(), time.Now(), "Bruce", "Wayne", "951-123-4567", 1,1, "A213467", "notbatman@wayne.com", 1, ""}
    //invalidUser := &User{uuid.Nil, time.Now(), time.Now(), "Thomas", "Wayne", "951-123-4567", 1,1, "", "notbatman@wayne.", -1, ""}
    fmt.Println("============")
    err := CreateUserEntry(validUser)

    ms.NoError(err)

    getUser := &User{}
    findErr := ms.DB.Find(getUser, "notbatman@wayne.com")
    ms.NoError(findErr)


}

func (ms *ModelSuite) Test_GetUser() {
    validUser := &User{}
    err := CreateUserEntry(validUser)
    ms.NoError(err)
    //t.Fatal("This test needs to be implemented!")
}

