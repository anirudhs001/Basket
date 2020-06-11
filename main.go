package main

import (
	"net/http"

	"./route"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", route.Index)
	http.HandleFunc("/customerPage", route.CustomerPage)
	http.HandleFunc("/sellerPage", route.SellerPage)

	http.ListenAndServe(":8080", nil)
}
