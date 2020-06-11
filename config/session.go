package config

import (
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

//GetCookie returns a cookie
func GetCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	c, err := req.Cookie("user_info")
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
	c := GetCookie(w, req)

	c.Value = c.Value + "|" + uType + "|" + fName + "|" + uName
	http.SetCookie(w, c)
}

//GetUser returns the userdata
func GetUser(w http.ResponseWriter, req *http.Request) (string, string, string) {

	c := GetCookie(w, req)

	temp := c.Value
	l := strings.Split(temp, "|")
	uType := l[1]
	fName := l[2]
	uName := l[3]
	return uType, fName, uName
}
