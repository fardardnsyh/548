package sqlFuncs

import (
	"database/sql"
	"fmt"
)

func PrintSql() {
	fmt.Println("hello")
}

func InitDatabase(dbLocation string) (*sql.DB, error) {
	return sql.Open("sqlite3", "./static/forum.db")
}

// indexPostStart     string = "<script>function %s() { var x = document.getElementById(\"%s\"); if (x.style.display === \"none\") { x.style.display = \"block\"; } else { x.style.display = \"none\";}} </script><div class=\"posts\"><div class=\"post\"><div class=\"title-wrapper\"><post-title>%s</post-title></div><div class=\"author-category-wrapper\"><div class=\"user-info\"><img src=\"../static/assets/middle-panel/user-icon.png\" alt=\"\"><div class=\"username-date\"><user-title>%s</user-title><post-date-title>%s</post-date-title></div></div>"
// 	// then we print out the categorys of that post see notes
// 	indexPostCategorysStart string = "<div class=\"categories\">"
// 	//uses only one trailing var(but uses it twice) cetegory wbhich is split from the categorys field in the post table
// 	indexPostCategorys    string = "<div class=\"category %s\">%s</div>"
// 	indexPostCategorysEnd string = "</div>"
// 	// use Fprintf with the trailing vars in this order post description | "?postid=" + postid for user nothing for guest | post likes | "?postid=" + postid for user nothing for guest | post dislikes | funcname(same as indexPostStart) | post num of coments
// 	indexPostEnd string = "</div><div class=\"body\"><post-text>%s</post-text></div><div class=\"interactions\"><div class=\"interaction-wrapper\"><div class=\"btn-dashboard\"><a href=\"/likePost%s\"><img src=\"../static/assets/middle-panel/thumbs-up.svg\" alt=\"\"></a></div><interaction-count>%s</interaction-count></div><div class=\"interaction-wrapper\"><div class=\"btn-dashboard\"><a href=\"/dislikePost%s\"><img src=\"../static/assets/middle-panel/thumbs-down.svg\" alt=\"\"></a></div><interaction-count>%s</interaction-count></div><div class=\"interaction-wrapper\"><button onclick=\"%s()\" class=\"toggle-comments\"><div class=\"btn-dashboard\">                                                <a href=\"#\"><img src=\"../static/assets/middle-panel/comments.svg\" alt=\"\"></a></div></button><interaction-count>%s</interaction-count></div></div></div>"
// 	// use Fprinf with the trailing var elementID(same as indexPostStart)
// 	indexCommentPostStart string = "<div id=\"%s\"><div class=\"title\">Comments</div>"
// 	// trailing var is post id
// 	indexCommentForm string = "<div class=\"create-comment\"><form action=\"/crecom\"><input type=\"text\" name=\"comment\" placeholder=\"Post a comment...\"><input type=\"hidden\" id=\"postid\" name=\"postid\" value=\"%s\"><input type=\"submit\" value=\"Post Comment\"></form></div>"
// 	// indexCommentPostStart should only be used at the start of the comments
// 	indexCommentWrapperStart string = "<div class=\"comments-wrapper\">"
// 	// use Fprinf with the trailing vars commentCreator | commentCreationDate | CommentBody | "?comid=" + comid for users nothing for guests | CommentLikes | "?comid=" + comid for users nothing for guests | CommentDislikes
// 	indexCommentsStart string = "<div class=\"comment\"><div class=\"user-info\"><img src=\"../static/assets/middle-panel/user-icon.png\" alt=\"\"><div class=\"username-date\"><user-title>%s</user-title><post-date-title>%s</post-date-title></div><img src=\"../static/assets/middle-panel/arrow.svg\" alt=\"\"></div><div class=\"comment-wrapper\"><div class=\"comment-text\">“%s”</div><div class=\"interactions\"><div class=\"interaction-wrapper\"><div class=\"btn-dashboard\"><a href=\"/comlike%s\"><img src=\"../static/assets/middle-panel/thumbs-up.svg\" width=\"28px\" alt=\"\"></a></div><interaction-count>%s</interaction-count></div><div class=\"interaction-wrapper\"><div class=\"btn-dashboard\"><a href=\"/comdislike%s\"><img src=\"../static/assets/middle-panel/thumbs-down.svg\" width=\"28px\" alt=\"\"></a></div><interaction-count>%s</interaction-count></div></div></div></div></div>"
// 	// indexCommentEnd should only be used when all comments have been printed out
// 	indexCommentEnd      string = "</div></div></div>"
// 	indexRightPanelStart string = "<div class=\"right-panel\"><div class=\"btn-text\">Registered Users</div><div class=\"registered-users\">"
// 	// use Fprintf with the trailing var username
// 	indexRightPanelUsers string = "<div class=\"user\"><img src=\"../static/assets/right-panel/user-icon.png\" alt=\"User Icon Image\"><user-title-right>%s</user-title-right></div>"
// 	indexEnd             string = "</div></div></div></div></div></body></html>"

// func home2(w http.ResponseWriter, r *http.Request) {
// 	var HTMLStrg1 string     /*This is the HTML that is sent to the http.ResponseWriter for it to display.*/
// 	var HTMLStrg2 string     /*This is the HTML that is sent to the http.ResponseWriter for it to display.*/
// 	var HTMLStrg3 string     /*This is the HTML that is sent to the http.ResponseWriter for it to display.*/
// 	var totalHTMLStrg string /*totalHTMLStrg = HTMLStrg1 + HTMLStrg2 + HTMLStrg3*/
// 	var idx int = 0
// 	var numOfDisplayPosts int = 0
// 	var numOfUserNames int = 0

// 	fmt.Println("Home2 → Main page handler entered again.")

// 	fmt.Fprint(w, `
// 	<html lang="en">
// 	<head>
// 		<meta charset="UTF-8">
// 		<meta http-equiv="X-UA-Compatible" content="IE=edge">
// 		<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 		<link rel="stylesheet" href="../static/css/styles.css">
// 		<script src="../static/js/scripts.js"></script>
// 		<link rel="shortcut icon" href="../static/assets/favicon.png" type="image/x-icon">
// 		<title>Forum | Home</title>
// 	</head>
// 	<body>
// 		<!-- the wrapper is to ensure everything is centered inside (flex) -->
// 		<div class="wrapper">
// 			<!-- seperated into four sections: navbar, left panel, mid section & right panel  -->
// 			<div class="content">

// 				<!-- NAVBAR -->
// 				<!-- this is to ensure everything is centered inside -->
// 				<div class="nav-wrapper">
// 					<!-- seperated into two sections: logo (left) & button (right) -->
// 					<div class="nav-content">
// 						<!-- this will be on the left side -->
// 						<div class="logo">
// 							<a href="/">
// 								<img src="../static/assets/navbar/logo.svg">
// 							</a>
// 						</div>
// 	`)

// 	if usrLoggedIn == 0 {
// 		fmt.Fprint(w, `
// 		<div class="btn-login">
// 							<a href="login">
// 								<img src="../static/assets/navbar/btn-login.svg" alt="login button">
// 							</a>
// 						</div>
// 		`)
// 	} else {
// 		HTMLStrg1 = `<div class="welcome-text">Welcome, `
// 		HTMLStrg1 = HTMLStrg1 + utilities.DBInfo.SignedInUser
// 		HTMLStrg1 = HTMLStrg1 + ` !</div><div class="btn-login">
// 					<a href="logout"><img src="../static/assets/navbar/btn-logout.svg" alt="logout button"></a></div>`
// 		fmt.Fprint(w, HTMLStrg1)
// 	} /*if there is a User Signed In*/

// 	fmt.Fprint(w, `
// 	</div>
// 				</div>

// 				<div class="sections">

// 					<!-- LEFT PANEL -->
// 					<div class="left-panel">

// 						<!-- dashboard btn -->
// 						<div class="btn dashboard">
// 							<div class="btn-dashboard">
// 								<a href="/">
// 									<img src="../static/assets/left-panel/btn-dashboard.svg" alt="Image of Home Page">
// 								</a>
// 							</div>
// 							<a href="/" class="btn-text">Dashboard</a>
// 						</div>

// 						<!-- filter by posts -->
// 						<div class="filter-posts">
// 							<div class="btn-text">Filter by Posts&nbsp&nbsp^</div>
// 						</div>

// 						<!-- your posts -->
// 						<div class="btn">
// 							<div class="btn-dashboard">
// 								<a href="#">
// 									<img src="../static/assets/left-panel/btn-your-posts.svg" alt="">
// 								</a>
// 							</div>
// 							<a href="yourPosts" class="btn-text">Your Posts</a>
// 						</div>

// 						<!-- your liked -->
// 						<div class="btn">
// 							<div class="btn-dashboard">
// 								<a href="#">
// 									<img src="../static/assets/left-panel/btn-liked-posts.svg" alt="">
// 								</a>
// 							</div>
// 							<a href="#" class="btn-text">Your Liked Posts</a>
// 						</div>

// 						<!-- filter by category -->
// 						<div class="filter-posts">
// 							<div class="btn-text">Filter by Category&nbsp&nbsp^</div>
// 						</div>

// 						<!-- categories -->
// 						<!-- to format each button horizontally using display: flex -->
// 						<div class="categories">
// 						<div class="category sports">`+utilities.DBInfo.CatName1+`</div>
// 						<div class="category food">`+utilities.DBInfo.CatName2+`</div>
// 						<div class="category technology">`+utilities.DBInfo.CatName3+`</div>
// 							<div class="category travel">`+utilities.DBInfo.CatName4+`</div>
// 							<div class="category miscellaneous">`+utilities.DBInfo.CatName5+`</div>
// 						</div>

// 					</div>

// 					<!-- MID SECTION -->
// 					<div class="mid-section">
// 						<!-- wrapper used for centering info-box -->
// 						<div class="info-wrapper">

// 							<!-- <div class="info-box">
// 								<info-text>In order to interact with posts, please login or create an account <a href="#">here.</a></info-text>
// 							</div> -->

// 						<form action="postcreate" name="postFrm" onsubmit="getChkBoxValue()">

// 							<!-- Create A Post/Title -->
// 							<div class="info-box">
// 								<input name="title" id="title" type="text" placeholder="Post title...">
// 								<div class="plus">
// 									<img src="../static/assets/middle-panel/plus.svg" alt="">
// 								</div>
// 							</div>

// 							<!-- Post Description -->
// 							<div class="info-box">
// 								<textarea placeholder="Post description..." id="description" name="description" rows="4" cols="50"></textarea>
// 							</div>

// 							<!-- Post Category -->
// 							<div class="info-box">
// 								<div class="container">
// 								  <div class="category sports">
// 									`+utilities.DBInfo.CatName1+`
// 									<input type="checkbox" style="display:inline" name="chkbox3" id="chkbox3" onclick="getChkBox3Value()" value="Movies"><span></span>
// 								  </div>
// 								  <div class="category food">
// 								  `+utilities.DBInfo.CatName2+`
// 								  <input type="checkbox" style="display:inline" name="chkbox2" id="chkbox2" onclick="getChkBox2Value()" value="Food"><span></span>
// 								  </div>
// 								  <div class="category technology">
// 								  `+utilities.DBInfo.CatName3+`
// 										<input type="checkbox" style="display:inline;" name="chkbox1" id="chkbox1" onclick="getChkBox1Value()" value="Technology"><span></span>
// 								  </div>
// 									<div class="category travel">
// 										`+utilities.DBInfo.CatName4+`
// 										<input type="checkbox" style="display:inline" name="chkbox4" id="chkbox4" onclick="getChkBox4Value()" value="Games"><span></span>
// 									</div>
// 									<div class="category miscellaneous">
// 										`+utilities.DBInfo.CatName5+`
// 										<input type="checkbox" style="display:inline" name="chkbox5" id="chkbox5" onclick="getChkBox5Value()" value="Miscellaneous"><span></span>
// 									</div>
// 								</div>
// 							</div>

// 							<!-- Create Post -->
// 							<div class="info-box">
// 								<input type="submit" value="Create Post" title="Click here to submit your Post !" alt="Submit Post button">
// 								<img src="../static/assets/middle-panel/tick.svg" title="Click here to submit your Post !" alt="Submit Post button">
// 							</div>

// 						</form>
// 						</div>

// 						<div class="posts-wrapper">
// 	`)
// 	/*Go code to Display the Posts information, display each and every post, Post info is obtained from the Global variable: displayPost.*/

// 	numOfDisplayPosts = len(utilities.DisplayPost)
// 	for idx = 0; idx < numOfDisplayPosts; idx++ {
// 		HTMLStrg1 = `
// 		<div class="posts">
// 		<div class="post">
// 			<!-- title-wrapper centers post-title -->
// 			<div class="title-wrapper">
// 				<post-title>` + utilities.DisplayPost[idx].Title + `</post-title>
// 			</div>

// 			<div class="author-category-wrapper">
// 				<!-- profile pic, username, post timestamp -->
// 				<div class="user-info">
// 					<img src="../static/assets/middle-panel/user-icon.png" alt="">
// 					<div class="username-date">
// 						<user-title>` + utilities.DisplayPost[idx].Username + `</user-title>
// 						<post-date-title>` + utilities.DisplayPost[idx].Agoduration + `</post-date-title>
// 					</div>
// 				</div>`

// 		HTMLStrg2 = `<!--this centers each category --><div class="categories">`
// 		if utilities.DisplayPost[idx].Category1 == "" {
// 		} else {
// 			HTMLStrg2 = HTMLStrg2 + `<div class="category food">` + utilities.DisplayPost[idx].Category1 + `</div>`
// 		} /*if Category empty*/
// 		if utilities.DisplayPost[idx].Category2 == "" {
// 		} else {
// 			HTMLStrg2 = HTMLStrg2 + `<div class="category technology">` + utilities.DisplayPost[idx].Category2 + `</div>`
// 		} /*if Category empty*/
// 		if utilities.DisplayPost[idx].Category3 == "" {
// 		} else {
// 			HTMLStrg2 = HTMLStrg2 + `<div class="category miscellaneous">` + utilities.DisplayPost[idx].Category3 + `</div>`
// 		} /*if Category empty*/
// 		HTMLStrg2 = HTMLStrg2 + `</div></div>`

// 		HTMLStrg3 = `<div class="body">
// 				<post-text>` + utilities.DisplayPost[idx].Body + `</post-text>
// 			</div>

// 			<div class="interactions">
// 				<!-- thumbs up - no of Likes -->
// 				<div class="interaction-wrapper">
// 					<div class="btn-dashboard">
// 						<a href="#">
// 							<img src="../static/assets/middle-panel/thumbs-up.svg" alt="">
// 						</a>
// 					</div>
// 					<interaction-count>` + utilities.DisplayPost[idx].NumOfLikes + `</interaction-count>
// 				</div>
// 				<!-- thumbs down - no of DisLikes -->
// 				<div class="interaction-wrapper">
// 					<div class="btn-dashboard">
// 						<a href="#">
// 							<img src="../static/assets/middle-panel/thumbs-down.svg" alt="">
// 						</a>
// 					</div>
// 					<interaction-count>` + utilities.DisplayPost[idx].NumOfDisLikes + `</interaction-count>
// 				</div>
// 				<!-- number of comments -->
// 				<div class="interaction-wrapper">
// 					<button onclick="myFunction()" class="toggle-comments">
// 						<div class="btn-dashboard">
// 							<a href="#">
// 								<img src="../static/assets/middle-panel/comments.svg" alt="">
// 							</a>
// 						</div>
// 					</button>
// 					<interaction-count>` + utilities.DisplayPost[idx].NumOfComents + `</interaction-count>
// 				</div>
// 			</div>
// 		</div>
// 		<!-- Post Comments -->
// 		<div id="comments">
// 			<div class="title">Comments</div>

// 			<div class="create-comment">
// 				<input type="text" name="comment" placeholder="Post a comment...">
// 				<input type="submit" value="Post Comment">
// 			</div>

// 			<div class="comments-wrapper">
// 				<!-- sample comment -->
// 				<div class="comment">
// 					<div class="user-info">
// 						<img src="../static/assets/middle-panel/user-icon.png" alt="">
// 						<div class="username-date">
// 							<user-title>` + utilities.DisplayPost[idx].Username + `</user-title>
// 							<post-date-title>01-01-01</post-date-title>
// 						</div>
// 						<img src="../static/assets/middle-panel/arrow.svg" alt="">
// 					</div>
// 					<div class="comment-wrapper">
// 						<div class="comment-text">` + utilities.DisplayPost[idx].Title + `</div>
// 						<div class="interactions">
// 							<!-- Comment Likes -->
// 							<div class="interaction-wrapper">
// 								<div class="btn-dashboard">
// 									<a href="#">
// 										<img src="../static/assets/middle-panel/thumbs-up.svg" width="28px" alt="">
// 									</a>
// 								</div>
// 								<interaction-count>` + utilities.DisplayPost[idx].NumOfLikes + `</interaction-count>
// 							</div>
// 							<!-- Comment Dislikes -->
// 							<div class="interaction-wrapper">
// 								<div class="btn-dashboard">
// 									<a href="#">
// 										<img src="../static/assets/middle-panel/thumbs-down.svg" width="28px" alt="">
// 									</a>
// 								</div>
// 								<interaction-count>` + utilities.DisplayPost[idx].NumOfDisLikes + `</interaction-count>
// 							</div>
// 						</div>
// 					</div>
// 				</div>
// 			</div>
// 		</div>
// 		</div>`
// 		totalHTMLStrg = HTMLStrg1 + HTMLStrg2 + HTMLStrg3
// 		fmt.Fprint(w, totalHTMLStrg)
// 	} /*for loop*/

// 	fmt.Fprint(w, `
// 					</div>
// 					</div>

// 					<!-- RIGHT PANEL -->
// 					<div class="right-panel">
// 						<div class="btn-text">Registered Users</div>
// 						<div class="registered-users">
// 	`)

// 	/*Go code to Display the User Names, display all Registered User names, User Names are obtained from the Global variable: DBInfo → User Names.*/
// 	numOfUserNames = len(utilities.DBInfo.UserName)
// 	for idx = 0; idx < numOfUserNames; idx++ {
// 		fmt.Fprint(w, `
// 					<div class="user">
// 						<img src="../static/assets/right-panel/user-icon.png" alt="User Icon Image">
// 						<user-title-right>`+utilities.DBInfo.UserName[idx]+`</user-title-right>
// 					</div>
// 		`)
// 	} /*for loop*/

// 	fmt.Fprint(w, `</div></div></div></div></div></body></html>`)
// } /*home2*/
