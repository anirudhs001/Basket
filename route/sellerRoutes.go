package route

import (
	"fmt"
	"net/http"

	"../config"
	"../models"
)

//SellerPage : main seller page
func SellerPage(w http.ResponseWriter, req *http.Request) {

	if _, userExists, _ := config.UserExists(req); !userExists {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//get current user(seller) details
	userType, currUser := config.GetUser(w, req)
	if userType != "seller" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	//TODO: send shop details and all orders
	//1) get shop details
	shopDetails, err := models.ReadSellerDetailsITEMSDB(currUser.Name)
	if err != nil {
		fmt.Println("SellerPage: err:")
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	//2) get all orders: DONE
	ordersList, err := models.ReadOrdersToSellerITEMSDB(currUser.Name)
	if err != nil {
		fmt.Println("SellerPage: err:")
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	type data struct {
		Shop models.Seller
		List []models.ShoppingList
	}

	//sanity check
	fmt.Println(shopDetails)
	fmt.Println(ordersList)
	config.Tpl.ExecuteTemplate(w, "sellerPage.gohtml", data{Shop: shopDetails, List: ordersList})
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
