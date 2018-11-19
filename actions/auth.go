// package actions

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gobuffalo/buffalo"
// 	"github.com/markbates/goth"
// 	"github.com/markbates/goth/gothic"
// 	"github.com/markbates/goth/providers/google"
// )

// func init() {
// 	gothic.Store = App().SessionStore

// 	goth.UseProviders(
// 		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
// 	)
// }

// func AuthCallback(c buffalo.Context) error {
// 	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
// 	if err != nil {
// 		return c.Error(401, err)
// 	}
// 	// Do something with the user, maybe register them/sign them in
// 	return c.Render(200, r.JSON(user))
// }
package actions

import (
	"net/http"
	"os"
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type lrnUser struct {
	user_fname string
	user_lname string
	user_phone string
	user_age string
	user_location string
	user_role string
	user_gender string
	user_id string
	user_email string
}

var (
	googleOauthConfig *oauth2.Config 
	user lrnUser
)

func (u lrnUser) printUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "name: %s\n", u.user_fname)
	fmt.Fprintf(w, "lname: %s\n", u.user_lname)
	fmt.Fprintf(w, "id: %s\n", u.user_id)
	fmt.Fprintf(w, "email: %s\n", u.user_email)
	fmt.Fprintf(w, "location: %s\n", u.user_location)
	fmt.Fprintf(w, "gender: %s\n", u.user_gender)
	fmt.Fprintf(w ,"phone: %s\n", u.user_phone)
	fmt.Fprintf(w, "age: %s\n", u.user_age)
}

func init() {
	// os.Setenv("CLIENT_ID", "190043352193-6gosbi41ard6f1itomqnd3u9kb831gtg.apps.googleusercontent.com")
 //    os.Setenv("SECRET_KEY", "mqkCA9p7dY1eej9TqmhkzFQx")
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://lrn-cse100/callback",
		ClientID: envy.Get("CLIENT_ID"),
		ClientSecret: envy.Get("SECRET_KEY"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
}

/*func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/process_new_user", handleSubmit)
	http.ListenAndServe(":8080", nil)
}*/

// func handleSubmit(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	user.user_location = r.Form["location"][0]
// 	user.user_gender = r.Form["gender"][0]
// 	user.user_age = r.Form["age"][0]
// 	user.user_role = r.Form["role"][0]
// 	user.user_phone = r.Form["phone"][0]
// 	user.printUserInfo(w, r)

// }

// func handleSignup(w http.ResponseWriter, r *http.Request) {
// 	var html = `<html>
// 				    <head>
// 				    <title></title>
// 				    </head>
// 				    <body>
// 				        <form action="/process_new_user" method="post">
// 				            Location:<input type="text" name="location">
// 				            Phone:<input type="text" name="phone">
// 				            Gender:<input type="text" name="gender">
// 				            Age:<input type="text" name="age">
// 				            <select name="role">
// 							  <option value="Tutor">Tutor</option>
// 							  <option value="Tutee">Tutee</option>
// 							</select>
// 				            <input type="submit" value="Submit">
// 				        </form>
// 				    </body>
// 				</html>`

// 	fmt.Fprint(w, html)
// }

// func handleHome(w http.ResponseWriter, r *http.Request) {
// 	var html = `<html>
// 					<body>
// 						<a href = "/login">Login</a></br>
// 						<a href = "/signup">Sign Up</a>
// 					</body>
// 				</html>`
// 	fmt.Fprint(w, html)
// }

func handleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	//read the oauth state from the cookie.
	oauthstate, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthstate.Value {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//THIS IS WHERE WE NEED TO STORE THE USER INFO AND THE OAUTHSTATE IN OUR DB,
	// AND REDIRECT TO ANOTHER PAGE.
	//WE NEED THE OAUTHSTATE TO LOG THE USER OUT.
	var guser map[string]interface{}
	jerr := json.Unmarshal([]byte(data), &guser)
	if jerr != nil {
		fmt.Printf("could not read json file: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	user.user_id = guser["id"].(string)
	user.user_fname = guser["given_name"].(string)
	user.user_lname = guser["family_name"].(string)
	user.user_email = guser["email"].(string)
	
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	//Let google authenticate the user and give us the info.

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %s\n", err.Error())
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("could not get request: %s\n", err.Error())
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse renponse: %s\n", err.Error())
	}

	return content, nil
}