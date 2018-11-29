package models

//package models_test

import (
    "time"

    "github.com/gobuffalo/uuid"
)

func (ms *ModelSuite) Test_User_Create() {
    count, err := ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(0, count)

    u := &User{
        ID:                 uuid.Nil,
        CreatedAt:          time.Now(),
        UpdatedAt:          time.Now(),
        FirstName:          "Bruce",
        LastName:           "Wayne",
        PhoneNumber:        "951-123-4567",
        GoogleID:           "ASD123",
        Email:              "bruce@wayne.com",
        Gender:             1,
        OtherSpecify:        "-",
    }
    verrs, err := u.Create(ms.DB)
    ms.NoError(err)
    ms.False(verrs.HasAny())

    count, err = ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
    count, err := ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(0, count)
    u := &User{
        FirstName: "bob",
    }

    verrs, err := u.Create(ms.DB)
    ms.NoError(err)
    ms.True(verrs.HasAny())

    count, err = ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
    count, err := ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(0, count)

    u := &User{
        ID:                 uuid.Nil,
        CreatedAt:          time.Now(),
        UpdatedAt:          time.Now(),
        FirstName:          "Bruce",
        LastName:           "Wayne",
        PhoneNumber:        "951-123-4567",
        GoogleID:           "ASD123",
        Email:              "bruce@wayne.com",
        Gender:             1,
        OtherSpecify:        "-",
    }
    verrs, err := u.Create(ms.DB)
    ms.NoError(err)
    ms.False(verrs.HasAny())

    count, err = ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(1, count)

    u = &User{
        ID:                 uuid.Nil,
        CreatedAt:          time.Now(),
        UpdatedAt:          time.Now(),
        FirstName:          "Thomas",
        LastName:           "Wayne",
        PhoneNumber:        "951-123-4567",
        GoogleID:           "ASD123",
        Email:              "bruce@wayne.com",
        Gender:             1,
        OtherSpecify:        "-",
    }
    verrs, err = u.Create(ms.DB)
    ms.Error(err)
    ms.False(verrs.HasAny())

    count, err = ms.DB.Count("users")
    ms.NoError(err)
    ms.Equal(1, count)
}
