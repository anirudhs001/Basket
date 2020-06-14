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
//returns all shops registered on website
func ViewSellers(w http.ResponseWriter, req *http.Request) {

	sellers, err := models.ReadAllSellersDB()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	config.Tpl.ExecuteTemplate(w, "viewSellers.gohtml", sellers)
}

//SellerPage : main seller page
func SellerPage(w http.ResponseWriter, req *http.Request) {

	config.Tpl.ExecuteTemplate(w, "sellerPage.gohtml", nil)
}

//SellerRegisterPage handler func
func SellerRegisterPage(w http.ResponseWriter, req *http.Request) {
	//TODO add login Page

	var seller models.Seller
	//store seller info
	if req.Method == http.MethodPost {
		seller = models.Seller{
			Name:      req.FormValue("name"),
			Addr:      req.FormValue("address"),
			OpenTime:  req.FormValue("opentime"),
			CloseTime: req.FormValue("closetime"),
		}
		if seller.Name == "" || seller.Addr == "" {
			fmt.Println("could not recieve seller data")
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		config.Tpl.ExecuteTemplate(w, "sellerRegisterPage.gohtml", seller)

		err := models.AddSellerDB(seller)
		if err != nil {
			if err == models.ErrSellerAlreadyExists {
				fmt.Println("seller already exists")
				fmt.Fprintf(w, "\nseller already exists!")
			} else {
				fmt.Println("ERROR occurred in SellerPage")
				fmt.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		} else {
			fmt.Println("shop added to database:" + seller.Name)
		}
	} else {
		config.Tpl.ExecuteTemplate(w, "sellerRegisterPage.gohtml", seller)
	}
}
