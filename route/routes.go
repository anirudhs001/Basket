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
	_, currUser := config.GetUser(w, req)

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
	config.Tpl.ExecuteTemplate(w, "customerPage.gohtml", shoppingList)
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
		fmt.Fprintf(w,
			`<li> %s - %s
			<button type="button" class="listButton" id="%s">del</button>
			</li>`,
			currUser.Name,
			newItem,
			id,
		)
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
func ViewSellers(w http.ResponseWriter, req *http.Request) {

	sellers, err := models.ReadSellersDB()
	if err != nil {
		fmt.Println(err)
	}
	config.Tpl.ExecuteTemplate(w, "viewSellers.gohtml", sellers)
}

//SellerPage handler func
func SellerPage(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
	//TODO make page
}
