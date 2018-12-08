package models


func (ms *ModelSuite) Test_Request_Create_Data() {
	count, err := ms.DB.Count("requests")
	ms.NoError(err)
	ms.Equal(0, count)

	u1 := &User{
		FirstName:    "Bruce",
		LastName:     "Wayne",
		PhoneNumber:  "951-123-4567",
		GoogleID:     "ASD123",
		Email:        "bruce@wayne.com",
		Gender:       2,
		OtherSpecify: "-",
        IsTutor:      1,
	}
	u2 := &User{
		FirstName:    "Thomas",
		LastName:     "Wayne",
		PhoneNumber:  "951-123-4567",
		GoogleID:     "ASD124",
		Email:        "thomas@wayne.com",
		Gender:       1,
		OtherSpecify: "-",
        IsTutor:      2,
	}
    _,_ = u1.Create(ms.DB)
    _,_ = u2.Create(ms.DB)
    //ms.NoError(err)
	//ms.False(verrs.HasAny())

	//count, err = ms.DB.Count("users")
	//ms.NoError(err)
	//ms.Equal(1, count)
    u1,_ = GetUserByGID(ms.DB, u1.GoogleID)
    u2,_ = GetUserByGID(ms.DB, u2.GoogleID)

    req, err := CreateNewRequestData(u1, u2)
    ms.NoError(err)
    ms.Equal(2, req.Status)
}
