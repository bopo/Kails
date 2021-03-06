package routes

import (
	"log"

	"github.com/go-martini/martini"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
)

// In this structure, the info
// to be sent will be stored
type FriendInfo struct {
	User   string
	Topics []string
	// Used to know if user is wheter connected or not
	isLogged bool
}

var (
	// Stores all user that has logged in/out
	// and their info
	allUsers []FriendInfo
	// Flag used to know if
	// a user has logged in/out
	hasStatusChanged bool
)

func Home(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Data["Title"] = "Home"
		ctx.Data["Name"] = ctx.User.Username
		ctx.Data["Language"] = ctx.User.StudyLanguage
		ctx.Data["Country"] = "Mexico"
		ctx.HTML(200, "user/home")
	} else {
		ctx.Data["Title"] = "Welcome"
		ctx.HTML(200, "main/main")
	}
}

func UserPage(ctx *middleware.Context, params martini.Params) {
	username := params["name"]
	user, err := models.UserByName(username)
	if err != nil {
		ctx.HTML(500, "")
		return
	}

	ctx.Data["Title"] = username
	ctx.Data["Username"] = user.Username
	ctx.Data["Email"] = user.Email
	ctx.Data["Since"] = user.Since.Format("January 2 of 2006")
	ctx.HTML(200, "user/info")
}

type SearchResult struct {
	Data  []string `json:"Data"`
	Error string   `json:"Error"`
}

func UserSearch(ctx *middleware.Context, params martini.Params) {
	var data SearchResult

	searchString := params["name"]
	results, err := models.UserSearch(searchString)

	if err != nil {
		data.Error = "couldn't find any users (with error)"
	}

	userList := make([]string, 0, 5)

	for _, user := range results {
		log.Println(user.Username)
		userList = append(userList, user.Username)
	}

	if len(userList) == 0 {
		data.Error = "couldn't find any users"
	} else {
		data.Data = userList
		data.Error = ""
	}

	ctx.JSON(200, data)

}

func Settings(ctx *middleware.Context) {
	ctx.Data["Title"] = "Settings"
	ctx.Data["currentLanguage"] = ctx.User.StudyLanguage
	var lang string
	if ctx.User.StudyLanguage == "english" {
		lang = "spanish"
	} else {
		lang = "english"
	}
	ctx.Data["otherLanguage"] = lang
	ctx.HTML(200, "user/settings")
}

type SettingsForm struct {
	Language string `form:"language"`
}

func SettingsPost(ctx *middleware.Context, form SettingsForm) {
	if form.Language == ctx.User.StudyLanguage {
		ctx.Redirect("/")
		return
	}

	err := ctx.User.UpdateStudyLanguage(form.Language)
	if err != nil {
		//TODO: put something here
	}

}

func SignUp(ctx *middleware.Context) {
	ctx.Data["Title"] = "Sign Up"
	ctx.HTML(200, "user/signup")
}

func SignUpPost(ctx *middleware.Context, form models.UserSignupForm) {
	err := models.NewUser(form)
	if err != nil {
		log.Println(err)
		ctx.HTML(501, "")
		return
	}
	ctx.Redirect("/")
}

func Login(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Redirect("/")
	} else {
		ctx.Data["Title"] = "Login"
		ctx.HTML(200, "user/login")
	}
}

func LoginPost(ctx *middleware.Context, form models.UserLoginForm) {
	var friend *FriendInfo
	friend = new(FriendInfo)

	ctx.Data["Title"] = "Home"
	user, err := models.UserByEmail(form.Email)
	if err != nil {
		// TODO: fill this up.
		log.Println(err)
		log.Println(user)
	}

	ctx.User = user
	ctx.Session.Set("name", user.Username)
	ctx.IsLogged = true

	// When someone logs in, the info is stores in the array allUsers
	friend.User = ctx.User.Username
	friend.Topics = ctx.User.Topics
	friend.isLogged = ctx.IsLogged
	allUsers = append(allUsers, *friend)

	// And the flag isJustConected is set to true
	hasStatusChanged = true

	ctx.Redirect("/")
}

func Logout(ctx *middleware.Context) {
	var friend *FriendInfo
	friend = new(FriendInfo)

	if ctx.IsLogged {
		// This is necessary
		ctx.Session.Clear()

		// This is for making sure
		ctx.User = new(models.User) // blank user
		ctx.IsLogged = false

		log.Println("user logged out")

		// The same when the user logs out
		friend.User = ctx.User.Username
		friend.Topics = ctx.User.Topics
		friend.isLogged = ctx.IsLogged
		allUsers = append(allUsers, *friend)

		hasStatusChanged = true
	}

	ctx.Redirect("/")
}

func UserInfo(ctx *middleware.Context, params martini.Params) {
	user, err := models.UserByName(params["user"])
	if err != nil {
		return
	}

	ctx.Data["Username"] = user.Username
	ctx.Data["Topics"] = user.Topics
	ctx.HTML(200, "user/user")
}

func Friends(ctx *middleware.Context) {
	ctx.Data["Title"] = "Friends"
	ctx.Data["Friends"], _ = ctx.User.ListFriends()
	ctx.HTML(200, "user/friends")
}
