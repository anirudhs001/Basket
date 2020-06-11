package route

import (
	"net/http"

	"../config"
	"../models"
)

//Index page
func Index(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet { //form sent via get
		userType := req.FormValue("userType")
		familyName := req.FormValue("familyName")
		userName := req.FormValue("userName")

		//TODO: set cookie: DONE
		config.SetUserType(w, req, userType, familyName, userName)

		if userType == "customer" {
			http.Redirect(w, req, "/customerPage", http.StatusTemporaryRedirect)
		}

		if userType == "seller" {
			http.Redirect(w, req, "/sellerPage", http.StatusTemporaryRedirect)
		}

	}
	//execute template
	config.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

//CustomerPage handlerfunc
func CustomerPage(w http.ResponseWriter, req *http.Request) {

	//get user data
	userType, familyName, userName := config.GetUser(w, req)

	currUser := models.User{
		parentGroup: familyName,
		name:        userName,
	}

	//list holding all items
	var List []models.User

	//query the family
	rows, err := db.Query("SELECT items from customers where name = $1", fname)
	if err != nil {
		http.NotFound(w, req)
	}

	var s string
	rows.Scan(&s)

	//get the updated items
	var new_item string
	if req.Method == http.MethodPost {
		new_item = req.FormValue("new_item")
	}

	//display the row:
	tpl.ExecuteTemplate(w, "familyPage.gohtml", nil)
}

func SellerPage(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
	//TODO make page
}
