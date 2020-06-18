package main

import (
	"net/http"

	"./route"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", route.Index)
	http.HandleFunc("/customerPage", route.CustomerPage)
	http.HandleFunc("/addItem", route.AddItem)
	http.HandleFunc("/delItem", route.DelItem)
	http.HandleFunc("/sellerPage", route.SellerPage)
	http.HandleFunc("/viewSellers", route.ViewSellers)
	http.HandleFunc("/sellerRegisterPage", route.SellerRegisterPage)
	http.HandleFunc("/sendRequestToSeller", route.SendRequestToSeller)
	http.HandleFunc("/signOut", route.SignOut)
	//TODO: add favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//serve the scripts
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./public"))))
	//start tcp server on port 8080
	//TODO: change in production
	http.ListenAndServe(":8080", nil)
}
