package actions

func (as *ActionSuite) Test_Login() {
	//put a user into the session to simulate login
	as.Session().Set("current_user", "Shaya Parsa")
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), "Hello Shaya Parsa")
	as.Contains(res.Body.String(), "Logout")
}

func (as *ActionSuite) Test_Logout() {
	//session is empty so we need to see the login screen
	as.Session().Clear()
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), "Please Log in: ")
	as.Contains(res.Body.String(), "Google Login")
}
