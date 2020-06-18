package route

import (
	"fmt"
	"net/http"

	"../config"
	"../models"
)

//Index page
func Index(w http.ResponseWriter, req *http.Request) {

	var userType string
	_, userExists, _ := config.UserExists(req)

	//TODO :if user already exists, reroute to customer page :DONE
	if userExists == true {
		userType, _ = config.GetUser(w, req)
	}

	if req.Method == http.MethodPost { //form sent via post

		if userExists == false { //user does not exist; create user

			userType = req.FormValue("userType")
			familyName := req.FormValue("familyName")
			userName := req.FormValue("userName")
			//TODO: set cookie: DONE
			config.SetUserType(w, req, userType, familyName, userName)
		}
	}

	if userType == "customer" {
		fmt.Println("user redirected to customer Page")
		http.Redirect(w, req, "/customerPage", http.StatusTemporaryRedirect)
	}

	if userType == "seller" {
		fmt.Println("user redirected to seller Page")
		http.Redirect(w, req, "/sellerPage", http.StatusTemporaryRedirect)
	}
	//execute template
	config.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

//CustomerPage handlerfunc
func CustomerPage(w http.ResponseWriter, req *http.Request) {

	if _, userExists, _ := config.UserExists(req); !userExists {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//get user data
	userType, currUser := config.GetUser(w, req)
	if userType != "customer" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	//sanity check
	fmt.Println("user accesed page: " + currUser.Name)

	//list to hold all items
	var shoppingList []models.Item
	var err error
	//get the updated items
	shoppingList, err = models.ReadItemsDB(currUser.ParentGroup)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		fmt.Println(err)
	}
	//display items in page
	//add currUser data
	type st struct {
		User models.User
		List []models.Item
	}
	config.Tpl.ExecuteTemplate(w, "customerPage.gohtml", st{User: currUser, List: shoppingList})
}

//AddItem handles the requests to add items
func AddItem(w http.ResponseWriter, req *http.Request) {

	//get user
	_, currUser := config.GetUser(w, req)

	if req.Method == http.MethodPost {
		//read response from client
		newItem := req.FormValue("item")
		//add new item
		id, err := models.InsertItem(currUser.ParentGroup, newItem, currUser.Name)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		//sanity check
		fmt.Println("item:" + newItem + " added by " + currUser.Name)
		//update list of items
		req.Header.Set("Content-Type", "text/html;charset=utf-8")
		fmt.Fprintf(w, `<div class="row mx-auto mt-2" id="list"><div class="col-3 mx-auto" id="listItem-user">`+currUser.Name+`</div><div class="col-3 mx-auto" id="listItem-item">`+newItem+`</div><button class="btn btn-outline-danger btn-md list-button" type="button" id="`+id+`">Delete</button></div>`)
	}
}

//DelItem handles the delete item requests
func DelItem(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		_, currUser := config.GetUser(w, req)

		//TODO: delte item from db: DONE
		models.DeleteItem(currUser.ParentGroup, req.FormValue("itemID"))
		//send back response
		fmt.Fprintln(w, "item deleted")
	}
}

//ViewSellers handler func
//returns all shops registered on website
func ViewSellers(w http.ResponseWriter, req *http.Request) {

	if _, userExists, _ := config.UserExists(req); !userExists {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	userType, _ := config.GetUser(w, req)
	if userType != "customer" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	sellers, err := models.ReadAllSellersDB()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	config.Tpl.ExecuteTemplate(w, "viewSellers.gohtml", sellers)
}

//SendRequestToSeller sends the currUser request to seller
func SendRequestToSeller(w http.ResponseWriter, req *http.Request) {

	_, currUser := config.GetUser(w, req)

	if req.Method == http.MethodPost {

		sellerName := req.FormValue("sellerName")
		if sellerName == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			fmt.Println("SendItems: err: could not receive form value")
			return
		}

		err := models.SendRequestToSeller(currUser.ParentGroup, sellerName)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			fmt.Println("SendItems: err: ")
			fmt.Println(err)
			return
		}
		fmt.Fprintln(w, "request sent to seller!")
	}
}

//SignOut removes the user_info cookie and reroutes to index page
func SignOut(w http.ResponseWriter, req *http.Request) {

	err := config.DeleteUser(w, req)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

//ViewOrders shows all previous orders
func ViewOrders(w http.ResponseWriter, req *http.Request) {

	if _, userExists, _ := config.UserExists(req); !userExists {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	userType, currUser := config.GetUser(w, req)
	if userType != "customer" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	shoppingLists, err := models.ViewOrdersitemsDB(currUser.ParentGroup)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	config.Tpl.ExecuteTemplate(w, "viewOrders.gohtml", shoppingLists)
}
