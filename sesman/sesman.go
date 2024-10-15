package sesman

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/satori/uuid.go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	LoginUuid string
	UserName  string
	Email     string
	Password  string
}

// creates a password hash from the entered password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checks the supplied password against the password hash created when the user registered
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// creates the session for the logged on user by creating the cookie and adding the users info to the session table
func CreateSession(db *sql.DB, user User, w http.ResponseWriter, r *http.Request) string {
	var err error
	sesCookie := CreateSessionCookie()
	http.SetCookie(w, &sesCookie)
	w.WriteHeader(200)
	_, err = db.Exec("DELETE FROM Session where auth_uuid='" + user.LoginUuid + "'")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec("INSERT INTO Session (uuid,auth_uuid) VALUES ('" + sesCookie.Value + "','" + user.LoginUuid + "');")
	if err != nil {
		fmt.Println(err.Error())
	}
	return sesCookie.Value
}

// getUser usses the session uuid to get the users info from the users table
func GetUser(session_uuid string) (User, bool) {
	var user User
	db, err := sql.Open("sqlite3", "./static/forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if session_uuid != "" {
		err = db.QueryRow("SELECT auth_uuid FROM Session WHERE uuid = '" + session_uuid + "'").Scan(&user.LoginUuid)
		if err == sql.ErrNoRows {
			return user, false
		} else {
			err = db.QueryRow("SELECT * FROM Users WHERE uuid = ?", user.LoginUuid).Scan(&user.LoginUuid, &user.UserName, &user.Email, &user.Password)
			if err == sql.ErrNoRows {
				return user, false
			} else {
				return user, true
			}
		}
	} else {
		return user, false
	}
}

//Create session cookie holds the template for the cookie used to authenticate the users session
func CreateSessionCookie() http.Cookie {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	return http.Cookie{
		Name:   "sessionCookie",
		Value:  u1.String(),
		MaxAge: 2 * int(time.Hour),
	}
}

// checkSession checks to see if the user has logged in by checking the cookie and the session table if the cookie does not exist or the session table doesnt have a matching entry it returns false otherwise it returns true and returns the session uuid
func CheckSession(w http.ResponseWriter, r *http.Request, db *sql.DB) (bool, string) {
	var err error
	var cookie *http.Cookie
	cookie, err = r.Cookie("sessionCookie")
	if err != nil {
		fmt.Println(err)
		return false, ""
	} else {
		session, _ := db.Query("SELECT * FROM Session WHERE uuid = '" + cookie.Value + "'")
		defer session.Close()
		var id int
		var sessionUuid string
		var authUuid string
		count := 0
		for session.Next() {
			session.Scan(&id, &sessionUuid, &authUuid)
			count++
		}
		if count == 1 {
			return true, cookie.Value
		} else {
			return false, ""
		}
	}

}

// SeleteSession dletes the session from the session table without this entry the cookie is usless so the cookie is left alone and overwritten when the user logs in again
func DeleteSession(sesid string, db *sql.DB) {
	db.Exec("DELETE FROM Session where uuid='" + sesid + "'")
}
