package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"webFuncs/sesman"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/uuid.go"
)

// contains all the variables for the html content
var (
	db *sql.DB
	// in between the login start and end put the error message for failed authentication
	loginStart    string = "<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"../static/css/styles.css\"><script src=\"../static/js/chkBox.js\"></script><link rel=\"shortcut icon\" href=\"../static/assets/favicon.png\" type=\"image/x-icon\"><title>Forum | Login</title></head><body><div class=\"wrapper\"><div class=\"content\"><div class=\"nav-wrapper\"><div class=\"nav-content\"><div class=\"logo\"><a href=\"/\"><img src=\"../static/assets/navbar/logo.svg\"></a></div><div class=\"info-box\"><info-text>"
	loginEnd      string = "</info-text></div><div class=\"btn-login\"><a href=\"/register\"><img src=\"../static/assets/navbar/btn-register.svg\" title=\"Click here to Sign In\" alt=\"Image of a Login (Sign In) button\"></a></div></div></div><div class=\"background\"><div class=\"auth-left\"><img src=\"../static/assets/auth/discussion.svg\" alt=\"people having a discussion\"><h3>Ready to Join<br>The Discussion?</h3><p>Before you get started,<br>you’ll need to login.</p></div><div class=\"auth-right\"><img src=\"../static/assets/auth/fingerprint.svg\" alt=\"\"><h1>Authentication</h1><h2>Let's get you logged in.</h2><form action=\"/login\" method=\"post\"><input type=\"text\" name=\"username\" placeholder=\"Username/Email\"></input><input type=\"password\" name=\"password\" placeholder=\"Password\"></input><input type=\"submit\" name=\"Login\" placeholder=\"Login\"></input></form><p>Don’t have an account yet?</p><a href=\"/register\">Sign Up</a></div></div></div></div></body></html>"
	registerStart string = "<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"../static/css/styles.css\"><script src=\"../static/js/chkBox.js\"></script><link rel=\"shortcut icon\" href=\"../static/assets/favicon.png\" type=\"image/x-icon\"><title>Forum | Register</title></head><body><div class=\"wrapper\"><div class=\"content\"><div class=\"nav-wrapper\"><div class=\"nav-content\"><div class=\"logo\"><a href=\"/\"><img src=\"../static/assets/navbar/logo.svg\"></a></div><div class=\"info-box\"><info-text>"
	registerEnd   string = "</info-text></div><div class=\"btn-login\"><a href=\"/login\"><img src=\"../static/assets/navbar/btn-login.svg\" title=\"Click here to Sign In\" alt=\"Image of a Login (Sign In) button\"></a></div></div></div><div class=\"background\"><div class=\"auth-left\"><img src=\"../static/assets/auth/discussion.svg\" alt=\"people having a discussion\"><h3>Ready to Join<br>The Discussion?</h3><p>Before you get started,<br>you’ll need an account.</p></div><div class=\"auth-right\"><img src=\"../static/assets/auth/fingerprint.svg\" alt=\"\"><h1>Registration</h1><h2>Let's create you an account.</h2><form action=\"register\" method=\"post\"><input type=\"text\" name=\"username\" placeholder=\"Username\"></input><input type=\"email\" name=\"emailadr\" placeholder=\"Email Address\"></input><input type=\"password\" name=\"password\" placeholder=\"Password\"></input><input type=\"submit\" name=\"Register\" placeholder=\"Register\"></input></form><p>Already have an account?</p><a href=\"login\">Sign In</a></div></div></div></div></body></html>"
	errStart      string = "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\"><link rel=\"stylesheet\" href=\"/static/css/error.css\"><link rel=\"icon\" href=\"data:;base64,=\"><title>Document</title></head><body><div class=\"content\"><div class=\"back\"><div class=\"err\"><p><div>"
	errEnd        string = "</div></p><p><a href=\"/\" class=\"btn\">Go home</a></p></div></div></div></body></html>"
	indexStart    string = `
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="../static/css/styles.css">
		<script src="../static/js/scripts.js"></script>
		<link rel="shortcut icon" href="../static/assets/favicon.png" type="image/x-icon">
		<title>Forum | Home</title>
	</head>
	
	<body>
		<div class="wrapper">
			<div class="content">
				<div class="nav-wrapper">
					<div class="nav-content">
						<div class="logo"><a href="/"><img src="../static/assets/navbar/logo.svg"></a></div>`

	indexLogout string = `
						<div class="welcome-text">Welcome, %s</div>
						<div class="btn-login"><a href="/logout"><img src="../static/assets/navbar/btn-logout.svg" alt="logout button"></a></div>`

	indexLogin string = `
						<div class="btn-login"><a href="/login"><img src="../static/assets/navbar/btn-login.svg" alt="login button"></a></div>`

	indexLeftPanalStart string = `
					</div>
				</div>
				<div class="sections">
					<div class="left-panel">
						<div class="btn dashboard">
							<div class="btn-dashboard"><a href="/"><img src="../static/assets/left-panel/btn-dashboard.svg" alt="Image of Home Page"></a></div>
							<a href="/" class="btn-text">Dashboard</a>
						</div>
						<div class="filter-posts">`

	indexFilterByPost string = `
							<div class="btn-text">Filter by Posts&nbsp&nbsp^</div>
							<div class="btn">
								<div class="btn-dashboard"><a href="/yourposts"><img src="../static/assets/left-panel/btn-your-posts.svg" alt=""></a></div>
								<a href="/yourPosts" class="btn-text">Your Posts</a>
							</div>
							<div class="btn">
								<div class="btn-dashboard">
									<a href="/yourlikes">
										<img src="../static/assets/left-panel/btn-liked-posts.svg" alt="">
									</a>
								</div>
								<a href="/yourlikes" class="btn-text">Your Liked Posts</a>
							</div>`

	indexFilterByCat string = `
							<div class="filter-posts">
								<div class="btn-text">Filter by Category&nbsp&nbsp^</div>
							</div>
							<div class="categories">
								<div class="category movies"><a href="/filterpost?category=movies">Movies</a></div>
								<div class="category food"><a href="/filterpost?category=food">Food</a></div>
								<div class="category technology"><a href="/filterpost?category=technology">Technology</a></div>
								<div class="category games"><a href="/filterpost?category=games">games</a></div>
								<div class="category miscellaneous"><a href="/filterpost?category=miscellaneous">Miscellaneous</a></div>
							</div>
						</div>
					</div>`

	indexMidsectionStart string = `
					<div class="mid-section">`

	indexInfoWrapper string = `
						<div class="info-wrapper">`

	indexCreatePost string = `
							<form action="crepost" name="postFrm" onsubmit="getChkBoxValue()">
								<div class="info-box"><input name="title" id="title" type="text"
										placeholder="Post title...">
									<div class="plus"><img src="../static/assets/middle-panel/plus.svg" alt=""></div>
								</div>
								<div class="info-box"><textarea placeholder="Post description..." id="description"
										name="description" rows="4" cols="50"></textarea></div>
								<div class="info-box">
									<div class="container">
										<div class="category movies">Movies<input type="checkbox" style="display:inline;"
												name="chkbox1" id="chkbox1" onclick="getChkBox1Value()"
												value="Movies"><span></span></div>
										<div class="category food">Food<input type="checkbox" style="display:inline"
												name="chkbox2" id="chkbox2" onclick="getChkBox2Value()"
												value="Food"><span></span></div>
										<div class="category technology">Technology<input type="checkbox"
												style="display:inline" name="chkbox3" id="chkbox3"
												onclick="getChkBox3Value()" value="Technology"><span></span></div>
										<div class="category games">Games<input type="checkbox" style="display:inline"
												name="chkbox4" id="chkbox4" onclick="getChkBox4Value()"
												value="Games"><span></span></div>
										<div class="category miscellaneous">Miscellaneous<input type="checkbox"
												style="display:inline" name="chkbox5" id="chkbox5"
												onclick="getChkBox5Value()" value="Miscellaneous"><span></span></div>
									</div>
								</div>
								<div class="info-box"><input type="submit" value="Create Post"
										title="Click here to submit your Post !" alt="Submit Post button"><img
										src="../static/assets/middle-panel/tick.svg"
										title="Click here to submit your Post !" alt="Submit Post button"></div>
							</form>`

	indexStartPostWrapper string = `
						</div>
						<div class="posts-wrapper">`

	indexPostInfo string = `
							<script>function %s() { var x = document.getElementById("%s"); if (x.style.display === "none") { x.style.display = "block"; } else { x.style.display = "none"; } } </script>
							<div class="posts">
								<div class="post">
									<div class="title-wrapper">
										<post-title>%s</post-title>
									</div>
									<div class="author-category-wrapper">
										<div class="user-info"><img src="../static/assets/middle-panel/user-icon.png"
												alt="">
											<div class="username-date">
												<user-title>%s</user-title>
												<post-date-title>%s</post-date-title>
											</div>
										</div>
										<div class="categories">`

	indexCategoryPrint string = `									
											<div class="category %s">%s</div>`

	indexPostEnd string = `										
										</div>
									</div>
									<div class="body">
										<post-text>%s</post-text>
									</div>
									<div class="interactions">
										<div class="interaction-wrapper">
											<div class="btn-dashboard"><a href="%s"><img
														src="../static/assets/middle-panel/thumbs-up.svg" alt=""></a></div>
											<interaction-count>%s</interaction-count>
										</div>
										<div class="interaction-wrapper">
											<div class="btn-dashboard"><a href="%s"><img
														src="../static/assets/middle-panel/thumbs-down.svg" alt=""></a>
											</div>
											<interaction-count>%s</interaction-count>
										</div>
										<div class="interaction-wrapper"><button onclick="%s()" class="toggle-comments">
												<div class="btn-dashboard"> <a href="#"><img
															src="../static/assets/middle-panel/comments.svg" alt=""></a>
												</div>
											</button>
											<interaction-count>%s</interaction-count>
										</div>
									</div>
								</div>
								<div id="%s">
									<div class="title">Comments</div>`

	indexCreateComment string = `
									<div class="create-comment">
										<form action="/crecom">
											<input type="text" name="comment" placeholder="Post a comment...">
											<input type="hidden" id="postid" name="postid" value="%s"><input type="submit" value="Post Comment">
										</form>
									</div>`

	indexCommentWrapperStart string = `
									<div class="comments-wrapper">`

	indexComment string = `
										<div class="comment">
											<div class="user-info">
												<img src="../static/assets/middle-panel/user-icon.png" alt="">
												<div class="username-date">
													<user-title>%s</user-title>
													<post-date-title>%s</post-date-title>
												</div>
												<img src="../static/assets/middle-panel/arrow.svg" alt="">
											</div>
											<div class="comment-wrapper">
												<div class="comment-text">“%s”</div>
												<div class="interactions">
													<div class="interaction-wrapper">
														<div class="btn-dashboard">
															<a href="%s">
																<img src="../static/assets/middle-panel/thumbs-up.svg" width="28px" alt="">
															</a>
														</div>
														<interaction-count>%s</interaction-count>
													</div>
													<div class="interaction-wrapper">
														<div class="btn-dashboard">
															<a href="%s">
																<img src="../static/assets/middle-panel/thumbs-down.svg" width="28px" alt="">
															</a>
														</div>
														<interaction-count>%s</interaction-count>
													</div>
												</div>
											</div>
										</div>`

	indexCommentWrapperEnd string = `
									</div>`

	indexPostAndCommentEnd string = `
								</div>
							</div>`

	indexPostWrapperMidSectionEnd string = `
						</div>
					</div>`

	indexRightPanelStart string = `
					<div class="right-panel">
						<div class="btn-text">Registered Users</div>
						<div class="registered-users">`

	indexRegisteredUserList string = `
							<div class="user"><img src="../static/assets/right-panel/user-icon.png" alt="User Icon Image">
								<user-title-right>%s</user-title-right>
							</div>`

	indexEnd string = `
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
	
	</html>`
)

//prints the error page withe the supplied err mainly handles http errors
func PrintError(err string) string {
	return errStart + err + errEnd
}

// works the same as print error but works with the login html page
func printLogin(err string) (output string) {
	return loginStart + err + loginEnd
}

// authenticates the user and checks the password hash using bcrypt
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user sesman.User
	if r.Method == "GET" {
		fmt.Fprintln(w, printLogin(""))
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			fmt.Fprintln(w, printLogin("Username or Password not entered"))
		} else {
			session, _ := db.Query("SELECT * FROM Users WHERE username = '" + username + "'")
			defer session.Close()
			count := 0
			var passwordHash string
			for session.Next() {
				session.Scan(&user.LoginUuid, &user.UserName, &user.Email, &passwordHash)
				count++
			}
			if count == 1 {
				if sesman.CheckPasswordHash(password, passwordHash) {
					sesman.CreateSession(db, user, w, r)
					buildIndex(w, r, user)
				} else {
					fmt.Fprintln(w, printLogin("Username and password do not match"))
				}
			} else {
				fmt.Fprintln(w, printLogin("ERROR username not in database"))
			}
		}
	}
}

// prints the register page with the appropriate error string i.e. username taken
func printRegister(err string) (output string) {
	return registerStart + err + registerEnd
}

// checks and validates the register information
func checkRegisterInfo(user sesman.User) int {
	username, _ := db.Query("SELECT * FROM Users WHERE username = '" + user.UserName + "'")
	email, _ := db.Query("SELECT * FROM Users WHERE email = '" + user.Email + "'")
	userUUID, _ := db.Query(("SELECT * FROM Users WHERE uuid = '" + user.LoginUuid + "'"))
	for username.Next() {
		return 1
	}
	for email.Next() {
		return 2
	}
	for userUUID.Next() {
		return 3
	}
	if user.UserName == "" {
		return 4
	}
	if user.Email == "" {
		return 5
	}
	if user.Password == "" {
		return 6
	}
	if !strings.Contains(user.Email, "@") {
		return 7
	}
	return 0
}

// handles the registration page Fprints and checks return from checkRegisterInfo
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, printRegister(""))
	} else {
		var err error
		var user sesman.User
		user.LoginUuid = uuid.Must(uuid.NewV4(), err).String()
		user.UserName = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.Email = r.FormValue("emailadr")
	startagain:
		switch checkRegisterInfo(user) {
		case 1:
			fmt.Fprintln(w, printRegister("Username is taken please provide another"))
			return
		case 2:
			fmt.Fprintln(w, printRegister("Email has already been used please provide another"))
			return
		case 3:
			user.LoginUuid = uuid.Must(uuid.NewV4(), err).String()
			goto startagain
		case 4:
			fmt.Fprintln(w, printRegister("you must enter a Username"))
			return
		case 5, 7:
			fmt.Fprintln(w, printRegister("you must enter a valid email address "))
			return
		case 6:
			fmt.Fprintln(w, printRegister("you must enter a Password"))
			return
		}
		user.Password, _ = sesman.HashPassword(user.Password)
		db.Exec("INSERT INTO Users VALUES(?,?,?,?)", user.LoginUuid, user.UserName, user.Email, user.Password)
		sesman.CreateSession(db, user, w, r)
		buildIndex(w, r, user)
	}
}

// create post creates the post and validates the the info from the create post form
func createPost(w http.ResponseWriter, r *http.Request, user sesman.User) {
	var chkBoxes [5]string
	var checkBoxAll string = ""
	r.ParseForm()
	created := time.Now().Format("01-02-2006 15:04:05")
	postTitle := r.FormValue("title")
	postBody := r.FormValue("description")
	chkBoxes[0] = r.Form.Get("chkbox1")
	chkBoxes[1] = r.Form.Get("chkbox2")
	chkBoxes[2] = r.Form.Get("chkbox3")
	chkBoxes[3] = r.Form.Get("chkbox4")
	chkBoxes[4] = r.Form.Get("chkbox5")
	for _, cate := range chkBoxes {
		switch cate {
		case "Movies":
			checkBoxAll += "Movies" + "/"
		case "Food":
			checkBoxAll += "Food" + "/"
		case "Technology":
			checkBoxAll += "Technology" + "/"
		case "Games":
			checkBoxAll += "Games" + "/"
		case "Miscellaneous":
			checkBoxAll += "Miscellaneous" + "/"
		}
	}
	checkBoxAll = strings.TrimSuffix(checkBoxAll, "/")
	if postTitle != "" && postBody != "" && checkBoxAll != "" {
		_, err := db.Exec("INSERT INTO Posts (title,body,user_id,amount_likes,amount_dislikes,categorys,posted_on) VALUES ('" + postTitle + "','" + postBody + "','" + user.LoginUuid + "','" + "0" + "','" + "0" + "','" + checkBoxAll + "','" + created + "');")
		if err != nil {
			fmt.Println(err)
		}
	}
	buildIndex(w, r, user)
}

// createComment validates comment info and adds the info to the database
func createComment(w http.ResponseWriter, r *http.Request, user sesman.User) {
	comment := r.FormValue("comment")
	postid := r.FormValue("postid")
	if comment != "" {
		db.Exec("INSERT INTO Comments (comment,auth_id,post_id,commented_on)VALUES ('" + comment + "','" + user.LoginUuid + "','" + postid + "','" + time.Now().Format("01-02-2006 15:04:05") + "');")
	}
	buildIndex(w, r, user)
}

// like post adds likes to posts and makes sure the user hasnt liked or disliked the post before
func likePost(w http.ResponseWriter, r *http.Request, user sesman.User) {
	postID := r.FormValue("postid")
	countLikes := 0
	countDislikes := 0
	check, _ := db.Query("SELECT * FROM dislikes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + postID + ";")
	for check.Next() {
		countDislikes++
	}
	check.Close()
	check, _ = db.Query("SELECT * FROM likes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + postID + ";")
	defer check.Close()
	for check.Next() {
		countLikes++
	}
	if countLikes > 0 {
		buildIndex(w, r, user)
	} else if countDislikes > 0 {
		buildIndex(w, r, user)
	} else {
		likes := 0
		_, err := db.Exec("INSERT INTO likes (user_uuid, post_id) VALUES(?,?);", user.LoginUuid, postID)
		if err != nil {
			log.Fatal(err)
		}
		addLike, _ := db.Query("SELECT amount_likes FROM Posts WHERE id =" + postID + ";")
		defer addLike.Close()
		for addLike.Next() {
			addLike.Scan(&likes)
		}
		likes++
		db.Exec("UPDATE Posts SET amount_likes = " + fmt.Sprint(likes) + " WHERE id =" + postID + ";")
		buildIndex(w, r, user)
	}
}

// same as like post but with dislikes
func dislikePost(w http.ResponseWriter, r *http.Request, user sesman.User) {
	postID := r.FormValue("postid")
	count := 0
	countDislikes := 0
	check, _ := db.Query("SELECT * FROM dislikes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + postID + ";")
	for check.Next() {
		countDislikes++
	}
	check.Close()
	check, _ = db.Query("SELECT * FROM likes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + postID + ";")
	defer check.Close()
	for check.Next() {
		count++
	}
	if count > 0 {
		buildIndex(w, r, user)
	} else if countDislikes > 0 {
		buildIndex(w, r, user)
	} else {
		disLikes := 0
		db.Exec("INSERT INTO dislikes (user_uuid, post_id) VALUES (?,?);", user.LoginUuid, postID)
		addLike, _ := db.Query("SELECT amount_dislikes FROM Posts WHERE id =" + postID + ";")
		defer addLike.Close()
		for addLike.Next() {
			addLike.Scan(&disLikes)
		}
		disLikes++
		db.Exec("UPDATE Posts SET amount_dislikes = " + fmt.Sprint(disLikes) + " WHERE id =" + postID + ";")
		buildIndex(w, r, user)
	}
}

// same as like post but with comments
func commentLike(w http.ResponseWriter, r *http.Request, user sesman.User) {
	comID := r.FormValue("comid")
	count := 0
	countDislikes := 0
	check, _ := db.Query("SELECT * FROM dislikes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + comID + ";")
	for check.Next() {
		countDislikes++
	}
	check.Close()
	check, _ = db.Query("SELECT * FROM likes WHERE user_uuid ='" + user.LoginUuid + "' AND com_id = '" + comID + "';")
	defer check.Close()
	for check.Next() {
		count++
	}
	if count > 0 {
		buildIndex(w, r, user)
	} else if countDislikes > 0 {
		buildIndex(w, r, user)
	} else {
		disLikes := 0
		db.Exec("INSERT INTO likes (user_uuid, com_id) VALUES (?,?);", user.LoginUuid, comID)
		addLike, _ := db.Query("SELECT amount_likes FROM Comments WHERE id =" + comID + ";")
		defer addLike.Close()
		for addLike.Next() {
			addLike.Scan(&disLikes)
		}
		disLikes++
		db.Exec("UPDATE Comments SET amount_likes = " + fmt.Sprint(disLikes) + " WHERE id =" + comID + ";")
		buildIndex(w, r, user)
	}
}

// same as dislike post but with somments
func commentDislike(w http.ResponseWriter, r *http.Request, user sesman.User) {
	comID := r.FormValue("comid")
	count := 0
	countDislikes := 0
	check, _ := db.Query("SELECT * FROM dislikes WHERE user_uuid ='" + user.LoginUuid + "' AND post_id = " + comID + ";")
	for check.Next() {
		countDislikes++
	}
	check.Close()
	check, _ = db.Query("SELECT * FROM likes WHERE user_uuid ='" + user.LoginUuid + "' AND com_id = '" + comID + "';")
	defer check.Close()
	for check.Next() {
		count++
	}
	if count > 0 {
		buildIndex(w, r, user)
	} else if countDislikes > 0 {
		buildIndex(w, r, user)
	} else {
		disLikes := 0
		db.Exec("INSERT INTO dislikes (user_uuid, com_id) VALUES (?,?);", user.LoginUuid, comID)
		addLike, _ := db.Query("SELECT amount_dislikes FROM Comments WHERE id =" + comID + ";")
		defer addLike.Close()
		for addLike.Next() {
			addLike.Scan(&disLikes)
		}
		disLikes++
		db.Exec("UPDATE Comments SET amount_dislikes = " + fmt.Sprint(disLikes) + " WHERE id =" + comID + ";")
		buildIndex(w, r, user)
	}
}

// logout dletes user from session table using DeleteSession and blanks the user struct
func logout(w http.ResponseWriter, r *http.Request, user sesman.User, sessionID string) {
	sesman.DeleteSession(sessionID, db)
	user.Email = ""
	user.LoginUuid = ""
	user.Password = ""
	user.UserName = ""
	buildIndex(w, r, user)
}

// builds the index page and only displays posts that the user has created
func yourPosts(w http.ResponseWriter, r *http.Request, user sesman.User) {
	fmt.Fprint(w, indexStart)
	loggenin := false
	fmt.Println(user.LoginUuid)
	if user.LoginUuid != "" {
		fmt.Fprintf(w, indexLogout, user.UserName)
		loggenin = true
	} else {
		fmt.Fprint(w, indexLogin)
	}
	fmt.Fprint(w, indexLeftPanalStart)
	if loggenin {
		fmt.Fprint(w, indexFilterByPost)
	}
	fmt.Fprint(w, indexFilterByCat+indexMidsectionStart+indexInfoWrapper)
	if loggenin {
		fmt.Fprint(w, indexCreatePost)
	}

	// ASC/DESC determines which order the posts are displayed on the homepage, with DESC meaning the latest posts are displayed first
	session, _ := db.Query("SELECT * FROM Posts WHERE user_id ='" + user.LoginUuid + "' ORDER BY posted_on DESC")
	defer session.Close()
	funcNameRune := 'A'
	elementIDRune := 'A'
	var id, amountLikes, amountDislikes, comId, postId, comlikes, comdislikes int
	var title, body, user_id, postedOn, userName, categorys, comment, auth_id, commentedOn string
	fmt.Fprint(w, indexStartPostWrapper)
	for session.Next() {
		session.Scan(&id, &title, &body, &user_id, &amountLikes, &amountDislikes, &categorys, &postedOn)
		category := strings.Split(categorys, "/")
		user, _ := db.Query("SELECT username from Users where uuid ='" + user_id + "';")
		for user.Next() {
			user.Scan(&userName)
		}
		fmt.Fprintf(w, indexPostInfo, "function"+string(funcNameRune), "elementID"+string(elementIDRune), title, userName, postedOn)
		for _, cat := range category {
			fmt.Fprintf(w, indexCategoryPrint, cat, cat)
		}
		if loggenin {
			fmt.Fprintf(w, indexPostEnd, body, "/likePost?postid="+fmt.Sprint(id), fmt.Sprint(amountLikes), "/dislikePost?postid="+fmt.Sprint(id), fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
			fmt.Fprintf(w, indexCreateComment, fmt.Sprint(id))
		} else {
			fmt.Fprintf(w, indexPostEnd, body, "", fmt.Sprint(amountLikes), "", fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
		}
		fmt.Fprint(w, indexCommentWrapperStart)
		comments, _ := db.Query("SELECT * from Comments where post_id =" + fmt.Sprint(id))
		for comments.Next() {
			var un string
			comments.Scan(&comId, &comment, &auth_id, &postId, &comlikes, &comdislikes, &commentedOn)
			user, _ := db.Query("SELECT username from Users where uuid ='" + auth_id + "';")
			for user.Next() {
				user.Scan(&un)
			}
			if loggenin {
				fmt.Fprintf(w, indexComment, un, commentedOn, comment, "/comlike?comid="+fmt.Sprint(comId), fmt.Sprint(comlikes), "/comdislike?comid="+fmt.Sprint(comId), fmt.Sprint(comdislikes))
			} else {
				fmt.Fprintf(w, indexComment, un, commentedOn, comment, "#", fmt.Sprint(comlikes), "#", fmt.Sprint(comdislikes))
			}
			user.Close()
		}
		//fmt.Fprint(w, indexCommentEnd)
		comments.Close()
		funcNameRune++
		elementIDRune++
		fmt.Fprint(w, indexCommentWrapperEnd)
		fmt.Fprint(w, indexPostAndCommentEnd)
	}

	fmt.Fprint(w, indexPostWrapperMidSectionEnd)
	fmt.Fprint(w, indexRightPanelStart)
	userList, _ := db.Query("SELECT username from Users")
	defer userList.Close()
	for userList.Next() {
		var un string
		userList.Scan(&un)
		fmt.Fprintf(w, indexRegisteredUserList, un)
	}
	fmt.Fprint(w, indexEnd)
}

// filterPost builda the index page and displays only the posts from the category that the user has clicked on
func filterPosts(w http.ResponseWriter, r *http.Request, user sesman.User) {
	r.ParseForm()
	selectedCat := r.FormValue("category")
	fmt.Fprint(w, indexStart)
	loggenin := false
	fmt.Println(user.LoginUuid)
	if user.LoginUuid != "" {
		fmt.Fprintf(w, indexLogout, user.UserName)
		loggenin = true
	} else {
		fmt.Fprint(w, indexLogin)
	}
	fmt.Fprint(w, indexLeftPanalStart)
	if loggenin {
		fmt.Fprint(w, indexFilterByPost)
	}
	fmt.Fprint(w, indexFilterByCat+indexMidsectionStart+indexInfoWrapper)
	if loggenin {
		fmt.Fprint(w, indexCreatePost)
	}
	session, _ := db.Query("SELECT * FROM Posts ORDER BY posted_on ASC")
	defer session.Close()
	funcNameRune := 'A'
	elementIDRune := 'A'
	var id, amountLikes, amountDislikes, comId, postId, comlikes, comdislikes int
	var title, body, user_id, postedOn, userName, categorys, comment, auth_id, commentedOn string
	fmt.Fprint(w, indexStartPostWrapper)
	for session.Next() {
		session.Scan(&id, &title, &body, &user_id, &amountLikes, &amountDislikes, &categorys, &postedOn)
		if strings.Contains(categorys, selectedCat) {
			category := strings.Split(categorys, "/")
			user, _ := db.Query("SELECT username from Users where uuid ='" + user_id + "';")
			for user.Next() {
				user.Scan(&userName)
			}
			fmt.Fprintf(w, indexPostInfo, "function"+string(funcNameRune), "elementID"+string(elementIDRune), title, userName, postedOn)
			for _, cat := range category {
				fmt.Fprintf(w, indexCategoryPrint, cat, cat)
			}
			if loggenin {
				fmt.Fprintf(w, indexPostEnd, body, "/likePost?postid="+fmt.Sprint(id), fmt.Sprint(amountLikes), "/dislikePost?postid="+fmt.Sprint(id), fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
				fmt.Fprintf(w, indexCreateComment, fmt.Sprint(id))
			} else {
				fmt.Fprintf(w, indexPostEnd, body, "", fmt.Sprint(amountLikes), "", fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
			}
			fmt.Fprint(w, indexCommentWrapperStart)
			comments, _ := db.Query("SELECT * from Comments where post_id =" + fmt.Sprint(id))
			for comments.Next() {
				var un string
				comments.Scan(&comId, &comment, &auth_id, &postId, &comlikes, &comdislikes, &commentedOn)
				user, _ := db.Query("SELECT username from Users where uuid ='" + auth_id + "';")
				for user.Next() {
					user.Scan(&un)
				}
				if loggenin {
					fmt.Fprintf(w, indexComment, un, commentedOn, comment, "/comlike?comid="+fmt.Sprint(comId), fmt.Sprint(comlikes), "/comdislike?comid="+fmt.Sprint(comId), fmt.Sprint(comdislikes))
				} else {
					fmt.Fprintf(w, indexComment, un, commentedOn, comment, "#", fmt.Sprint(comlikes), "#", fmt.Sprint(comdislikes))
				}
				user.Close()
			}
			//fmt.Fprint(w, indexCommentEnd)
			comments.Close()
			funcNameRune++
			elementIDRune++
			fmt.Fprint(w, indexCommentWrapperEnd)
			fmt.Fprint(w, indexPostAndCommentEnd)
		}
	}

	fmt.Fprint(w, indexPostWrapperMidSectionEnd)
	fmt.Fprint(w, indexRightPanelStart)
	userList, _ := db.Query("SELECT username from Users")
	defer userList.Close()
	for userList.Next() {
		var un string
		userList.Scan(&un)
		fmt.Fprintf(w, indexRegisteredUserList, un)
	}
	fmt.Fprint(w, indexEnd)
}

// build index builds the vannilla version of the index page
func buildIndex(w http.ResponseWriter, r *http.Request, user sesman.User) {
	fmt.Fprint(w, indexStart)
	loggenin := false
	fmt.Println(user.LoginUuid)
	if user.LoginUuid != "" {
		fmt.Fprintf(w, indexLogout, user.UserName)
		loggenin = true
	} else {
		fmt.Fprint(w, indexLogin)
	}
	fmt.Fprint(w, indexLeftPanalStart)
	if loggenin {
		fmt.Fprint(w, indexFilterByPost)
	}
	fmt.Fprint(w, indexFilterByCat+indexMidsectionStart+indexInfoWrapper)
	if loggenin {
		fmt.Fprint(w, indexCreatePost)
	}
	session, _ := db.Query("SELECT * FROM Posts ORDER BY posted_on ASC")
	defer session.Close()
	funcNameRune := 'A'
	elementIDRune := 'A'
	var id, amountLikes, amountDislikes, comId, postId, comlikes, comdislikes int
	var title, body, user_id, postedOn, userName, categorys, comment, auth_id, commentedOn string
	fmt.Fprint(w, indexStartPostWrapper)
	for session.Next() {
		session.Scan(&id, &title, &body, &user_id, &amountLikes, &amountDislikes, &categorys, &postedOn)
		category := strings.Split(categorys, "/")
		user, _ := db.Query("SELECT username from Users where uuid ='" + user_id + "';")
		for user.Next() {
			user.Scan(&userName)
		}
		fmt.Fprintf(w, indexPostInfo, "function"+string(funcNameRune), "elementID"+string(elementIDRune), title, userName, postedOn)
		for _, cat := range category {
			fmt.Fprintf(w, indexCategoryPrint, cat, cat)
		}
		if loggenin {
			fmt.Fprintf(w, indexPostEnd, body, "/likePost?postid="+fmt.Sprint(id), fmt.Sprint(amountLikes), "/dislikePost?postid="+fmt.Sprint(id), fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
			fmt.Fprintf(w, indexCreateComment, fmt.Sprint(id))
		} else {
			fmt.Fprintf(w, indexPostEnd, body, "", fmt.Sprint(amountLikes), "", fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
		}
		fmt.Fprint(w, indexCommentWrapperStart)
		comments, _ := db.Query("SELECT * from Comments where post_id =" + fmt.Sprint(id))
		for comments.Next() {
			var un string
			comments.Scan(&comId, &comment, &auth_id, &postId, &comlikes, &comdislikes, &commentedOn)
			user, _ := db.Query("SELECT username from Users where uuid ='" + auth_id + "';")
			for user.Next() {
				user.Scan(&un)
			}
			if loggenin {
				fmt.Fprintf(w, indexComment, un, commentedOn, comment, "/comlike?comid="+fmt.Sprint(comId), fmt.Sprint(comlikes), "/comdislike?comid="+fmt.Sprint(comId), fmt.Sprint(comdislikes))
			} else {
				fmt.Fprintf(w, indexComment, un, commentedOn, comment, "#", fmt.Sprint(comlikes), "#", fmt.Sprint(comdislikes))
			}
			user.Close()
		}
		//fmt.Fprint(w, indexCommentEnd)
		comments.Close()
		funcNameRune++
		elementIDRune++
		fmt.Fprint(w, indexCommentWrapperEnd)
		fmt.Fprint(w, indexPostAndCommentEnd)
	}

	fmt.Fprint(w, indexPostWrapperMidSectionEnd)
	fmt.Fprint(w, indexRightPanelStart)
	userList, _ := db.Query("SELECT username from Users")
	defer userList.Close()
	for userList.Next() {
		var un string
		userList.Scan(&un)
		fmt.Fprintf(w, indexRegisteredUserList, un)
	}
	fmt.Fprint(w, indexEnd)
}

// yourLikes builds the index page and only lists the posts you have liked
func yourLikes(w http.ResponseWriter, r *http.Request, user sesman.User) {
	var likedPostID int
	fmt.Fprint(w, indexStart)
	loggenin := false
	fmt.Println(user.LoginUuid)
	if user.LoginUuid != "" {
		fmt.Fprintf(w, indexLogout, user.UserName)
		loggenin = true
	} else {
		fmt.Fprint(w, indexLogin)
	}
	fmt.Fprint(w, indexLeftPanalStart)
	if loggenin {
		fmt.Fprint(w, indexFilterByPost)
	}
	fmt.Fprint(w, indexFilterByCat+indexMidsectionStart+indexInfoWrapper)
	if loggenin {
		fmt.Fprint(w, indexCreatePost)
	}
	liked, _ := db.Query("SELECT post_id FROM likes WHERE user_uuid ='" + user.LoginUuid + "'")
	defer liked.Close()
	for liked.Next() {
		liked.Scan(&likedPostID)
		session, _ := db.Query("SELECT * FROM Posts Where id =" + fmt.Sprint(likedPostID))
		defer session.Close()
		funcNameRune := 'A'
		elementIDRune := 'A'
		var id, amountLikes, amountDislikes, comId, postId, comlikes, comdislikes int
		var title, body, user_id, postedOn, userName, categorys, comment, auth_id, commentedOn string
		fmt.Fprint(w, indexStartPostWrapper)
		for session.Next() {
			session.Scan(&id, &title, &body, &user_id, &amountLikes, &amountDislikes, &categorys, &postedOn)
			category := strings.Split(categorys, "/")
			user, _ := db.Query("SELECT username from Users where uuid ='" + user_id + "';")
			for user.Next() {
				user.Scan(&userName)
			}
			fmt.Fprintf(w, indexPostInfo, "function"+string(funcNameRune), "elementID"+string(elementIDRune), title, userName, postedOn)
			for _, cat := range category {
				fmt.Fprintf(w, indexCategoryPrint, cat, cat)
			}
			if loggenin {
				fmt.Fprintf(w, indexPostEnd, body, "/likePost?postid="+fmt.Sprint(id), fmt.Sprint(amountLikes), "/dislikePost?postid="+fmt.Sprint(id), fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
				fmt.Fprintf(w, indexCreateComment, fmt.Sprint(id))
			} else {
				fmt.Fprintf(w, indexPostEnd, body, "", fmt.Sprint(amountLikes), "", fmt.Sprint(amountDislikes), "function"+string(funcNameRune), "", "elementID"+string(elementIDRune))
			}
			fmt.Fprint(w, indexCommentWrapperStart)
			comments, _ := db.Query("SELECT * from Comments where post_id =" + fmt.Sprint(id))
			for comments.Next() {
				var un string
				comments.Scan(&comId, &comment, &auth_id, &postId, &comlikes, &comdislikes, &commentedOn)
				user, _ := db.Query("SELECT username from Users where uuid ='" + auth_id + "';")
				for user.Next() {
					user.Scan(&un)
				}
				if loggenin {
					fmt.Fprintf(w, indexComment, un, commentedOn, comment, "/comlike?comid="+fmt.Sprint(comId), fmt.Sprint(comlikes), "/comdislike?comid="+fmt.Sprint(comId), fmt.Sprint(comdislikes))
				} else {
					fmt.Fprintf(w, indexComment, un, commentedOn, comment, "#", fmt.Sprint(comlikes), "#", fmt.Sprint(comdislikes))
				}
				user.Close()
			}
			//fmt.Fprint(w, indexCommentEnd)
			comments.Close()
			funcNameRune++
			elementIDRune++
			fmt.Fprint(w, indexCommentWrapperEnd)
			fmt.Fprint(w, indexPostAndCommentEnd)
		}
	}

	fmt.Fprint(w, indexPostWrapperMidSectionEnd)
	fmt.Fprint(w, indexRightPanelStart)
	userList, _ := db.Query("SELECT username from Users")
	defer userList.Close()
	for userList.Next() {
		var un string
		userList.Scan(&un)
		fmt.Fprintf(w, indexRegisteredUserList, un)
	}
	fmt.Fprint(w, indexEnd)
}

// the middleware validates the users session and user info if they are logged in and sends the user to the appropriate page also stops the user from accessing certain pages if the user is not logged in and same goes for logged in users
func middleware(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("sqlite3", "./static/forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	urlFrom := strings.TrimPrefix(r.URL.Path, "/")
	var user sesman.User
	var found bool
	if foundSession, sessionValue := sesman.CheckSession(w, r, db); foundSession {
		user, found = sesman.GetUser(sessionValue)
		if !found {
			buildIndex(w, r, user)
		}
		switch urlFrom {
		case "crepost":
			createPost(w, r, user)
		case "crecom":
			createComment(w, r, user)
		case "likePost":
			likePost(w, r, user)
		case "dislikePost":
			dislikePost(w, r, user)
		case "logout":
			logout(w, r, user, sessionValue)
		case "comlike":
			commentLike(w, r, user)
		case "comdislike":
			commentDislike(w, r, user)
		case "yourPosts":
			yourPosts(w, r, user)
		case "yourlikes":
			yourLikes(w, r, user)
		case "filterpost":
			filterPosts(w, r, user)
		case "":
			buildIndex(w, r, user)
		case "login":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "register":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		default:
			fmt.Fprintln(w, PrintError("404 not found"))
		}
	} else {
		switch urlFrom {
		case "crepost":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "crecom":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "likePost":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "dislikePost":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "logout":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "comlike":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "comdislike":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "yourPosts":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "yourlikes":
			fmt.Fprintln(w, PrintError("Method not allowed 405"))
		case "filterpost":
			filterPosts(w, r, user)
		case "login":
			loginHandler(w, r)
		case "register":
			registerHandler(w, r)
		case "":
			buildIndex(w, r, user)
		default:
			fmt.Fprintln(w, PrintError("404 not found"))
		}
	}
}

//all the create table sql statements are located in the init function so if the program loads and the tables do not exist creates them and it also sets up the database
func init() {
	var err error
	db, err = sql.Open("sqlite3", "./static/forum.db")
	if err != nil {
		fmt.Println(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		uuid TEXT PRIMARY KEY,
		username text NOT NULL unique CHECK(LENGTH(username) <= 40),
		email text NOT NULL unique,
		password_hash text NOT NULL
	  );`)
	db.Exec(`CREATE TABLE IF NOT EXISTS Session (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid text NOT NULL,
		auth_uuid TEXT,
		FOREIGN KEY (auth_uuid) REFERENCES user (uuid)
	  );`)
	db.Exec(`CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		user_id TEXT NOT NULL,
		amount_likes INTEGER NOT NULL,
		amount_dislikes INTEGER NOT NULL,
		categorys TEXT NOT NULL,
		posted_on TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user (uuid)
	  );`)
	db.Exec(`CREATE TABLE IF NOT EXISTS Comments (
		id integer PRIMARY KEY,
		comment text,
		auth_id TEXT NOT NULL,
		post_id integer NOT NULL,
		amount_likes INTEGER NOT NULL DEFAULT 0,
		amount_dislikes INTEGER NOT NULL DEFAULT 0,
		commented_on TEXT NOT NULL,
		FOREIGN KEY (auth_id) REFERENCES user (uuid) FOREIGN KEY (post_id) REFERENCES posts (id)
	  );`)
	db.Exec(`CREATE TABLE IF NOT EXISTS likes (
		id integer primary key autoincrement, 
		user_uuid text not null references Users (uuid), 
		post_id integer references Posts (id), 
		com_id integer references Comments (id));
	  `)
	db.Exec(`CREATE TABLE IF NOT EXISTS dislikes (
		id integer primary key autoincrement, 
		user_uuid text not null references Users (uuid), 
		post_id integer references Posts (id),
		com_id integer references Comments (id)
		);`)
}

// only the handler func for the middle ware and listen and serve are in the main function so the middle ware is used everytime so the users session can be validated
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", middleware)         // pointing the handler func to the middleware
	err := http.ListenAndServe(":8080", nil) // setting listening port and starting web server
	if err != nil {
		log.Fatalf(fmt.Sprintf("500 %s : %s\n", http.StatusText(500), err.Error()))
	}
}
