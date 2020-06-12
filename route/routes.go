package route

import (
	"fmt"
	"net/http"

	"../config"
	"../models"
)

//Index page
func Index(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost { //form sent via post
		userType := req.FormValue("userType")
		familyName := req.FormValue("familyName")
		userName := req.FormValue("userName")

		//TODO: set cookie: DONE
		config.SetUserType(w, req, userType, familyName, userName)

		if userType == "customer" {
			fmt.Println("user redirected to customer Page")
			http.Redirect(w, req, "/customerPage", http.StatusTemporaryRedirect)
		}

		if userType == "seller" {
			fmt.Println("user redirected to seller Page")
			http.Redirect(w, req, "/sellerPage", http.StatusTemporaryRedirect)
		}

	}
	//execute template
	config.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

//CustomerPage handlerfunc
func CustomerPage(w http.ResponseWriter, req *http.Request) {

	//get user data
	_, fname, uname := config.GetUser(w, req)

	currUser := models.User{
		ParentGroup: fname,
		Name:        uname,
	}
	//sanity check
	fmt.Println("user accesed page: " + currUser.Name)

	//list to hold all items
	var shoppingList []map[string]string
	var err error
	//get the updated items
	shoppingList, err = models.ReadItemsDB(w, req, currUser.ParentGroup)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		fmt.Println(err)
	}
	//display items in page
	config.Tpl.ExecuteTemplate(w, "customerPage.gohtml", shoppingList)
}

//AddItem handles the requests to add items
func AddItem(w http.ResponseWriter, req *http.Request) {

	//get user
	_, fname, uname := config.GetUser(w, req)
	currUser := models.User{
		ParentGroup: fname,
		Name:        uname,
	}

	if req.Method == http.MethodPost {
		//read response from client
		newItem := req.FormValue("item")
		//add new item
		err := models.InsertItem(w, req, currUser.ParentGroup, newItem, currUser.Name)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		//sanity check
		fmt.Println("item:" + newItem + " added by " + currUser.Name)
	}
}

//SellerPage handler func
func SellerPage(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
	//TODO make page
}
