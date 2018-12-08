package models

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &User{
		FirstName:    "Bruce",
		LastName:     "Wayne",
		PhoneNumber:  "951-123-4567",
		GoogleID:     "ASD123",
		Email:        "bruce@wayne.com",
		Gender:       1,
		OtherSpecify: "-",
        IsTutor:      1,
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
		FirstName:    "Bruce",
		LastName:     "Wayne",
		PhoneNumber:  "951-123-4567",
		GoogleID:     "ASD123",
		Email:        "bruce@wayne.com",
		Gender:       1,
		OtherSpecify: "-",
        IsTutor:      1,
	}
	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &User{
		FirstName:    "Thomas",
		LastName:     "Wayne",
		PhoneNumber:  "951-123-4567",
		GoogleID:     "ASD123",
		Email:        "bruce@wayne.com",
		Gender:       1,
		OtherSpecify: "-",
        IsTutor:      2,
	}
	verrs, err = u.Create(ms.DB)
	ms.Error(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
