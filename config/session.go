package config

import (
	"net/http"
	"strings"

	"../models"
	uuid "github.com/satori/go.uuid"
)

//GetCookie returns a cookie
func GetCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	c, _, err := UserExists(req)
	if err == http.ErrNoCookie { //create cookie
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "user_info",
			Value: sID.String(),
		}
	}
	if err != http.ErrNoCookie && err != nil {
		http.Error(w, "could not write cookie", http.StatusServiceUnavailable)
	}
	http.SetCookie(w, c)
	return c
}

//SetUserType sets the user type in cookie
func SetUserType(w http.ResponseWriter, req *http.Request, uType string, fName string, uName string) {

	_, err := req.Cookie("user_info")
	if err == http.ErrNoCookie {

		c := GetCookie(w, req)
		c.Value = c.Value + `|` + uType + `|` + fName + `|` + uName
		http.SetCookie(w, c)
	}

}

//GetUser returns the userdata
func GetUser(w http.ResponseWriter, req *http.Request) (string, models.User) {

	c := GetCookie(w, req)

	temp := c.Value
	l := strings.Split(temp, "|")
	uType := l[1]
	fName := l[2]
	uName := l[3]

	u := models.User{
		ParentGroup: fName,
		Name:        uName,
	}
	return uType, u
}

//UserExists returns true if user exists, false otherwise and the error, if any.
func UserExists(req *http.Request) (*http.Cookie, bool, error) {

	c, err := req.Cookie("user_info")
	if err != nil {
		return c, false, err
	}
	return c, true, nil
}

//DeleteUser deletes the user_info cookie(if exists). returns error if any
func DeleteUser(w http.ResponseWriter, req *http.Request) error {

	c, userExists, err := UserExists(req)
	if err != nil {
		return err
	}
	if userExists {
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
	return nil
}
