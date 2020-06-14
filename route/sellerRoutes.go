package route

import (
	"fmt"
	"net/http"

	"../config"
	"../models"
)

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
