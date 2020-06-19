package models

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	_ "github.com/lib/pq" //need the init function
)

var db *sql.DB //not exported

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://temp_user:password@localhost/basket?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("database connected")
}

//ReadItemsDB reads database and returns all the items
func ReadItemsDB(familyName string) ([]Item, error) {

	//query the family
	s, err := getCustomerRow(familyName)
	var items []Item
	if err != nil {
		return items, err
	}

	//read s if not empty
	if s != "" {
		li := strings.Split(s, "\n")
		for _, v := range li {
			if v != "" {

				id := strings.Split(v, "|")[0]
				user := strings.Split(v, "|")[1]
				item := strings.Split(v, "|")[2]
				items = append(items, Item{
					ID:   id,
					Name: user,
					Item: item,
				})
			}
		}
	}
	return items, err
}

//InsertItem appends the item to a familyrow in database
//returns the new items ID and nil if succesfull
//otherwise returns an emty string and the error if any
func InsertItem(familyName string, item string, user string) (string, error) {

	//query the row
	s, err := getCustomerRow(familyName)
	if err != nil {
		return "", err
	}

	//add new item
	b := make([]byte, 5)
	rand.Read(b)
	id := hex.EncodeToString(b)
	s = id + "|" + user + "|" + item + "\n" + s

	err = updateCustomerRow(s, familyName)
	if err != nil {
		return "", err
	}

	return id, err
}

//DeleteItem an item identified by its ID
func DeleteItem(familyName string, itemID string) error {

	s, err := getCustomerRow(familyName)
	if err != nil {
		return err
	}

	var sNew string //string to hold update items list
	if s != "" {
		li := strings.Split(s, "\n")
		for _, v := range li {
			if v != "" {
				id := strings.Split(v, "|")[0]
				if id == itemID { //delete this row
					continue
					//skip adding this row
				}
				sNew += v + "\n"
			}
		}
	}

	//write row to db
	err = updateCustomerRow(sNew, familyName)

	return err
}

//ReadAllSellersDB returns a list of sellers
func ReadAllSellersDB() ([]Seller, error) {

	var listOfSellers []Seller
	var s Seller

	rows, err := getSellerRows()
	if err != nil {
		return listOfSellers, err
	}

	for rows.Next() {
		if err = rows.Scan(&s.Name, &s.Addr, &s.OpenTime, &s.CloseTime); err != nil {
			return listOfSellers, err
		}
		listOfSellers = append(listOfSellers, s)
	}
	return listOfSellers, err
}

//AddSellerDB adds the seller to sellers table
//returns non-nil err if row already exists
func AddSellerDB(s Seller) error {

	//query to check seller row
	r := db.QueryRow("SELECT * FROM sellers where name=$1;", s.Name)

	//try to read row;
	//returns ErrNoRows if row does not exist
	var ts Seller
	err := r.Scan(&ts.Name, &ts.Addr, &ts.OpenTime, &ts.CloseTime)
	if err == sql.ErrNoRows {
		//create row
		err = insertSellerRow(s)
		return err
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	//else if err == nil
	return ErrSellerAlreadyExists
}

//SendRequestToSeller adds the users items list to "items" database
func SendRequestToSeller(familyName string, sellerName string) error {

	items, err := getCustomerRow(familyName)
	if err != nil {
		return err
	}

	//add items to db
	t := time.Now()
	// timeStamp := t.Format("2/Jan/2006 15:04")

	err = addRequestToitemsDB(t, familyName, sellerName, items)
	return err
}

//ReadOrdersByCustomerITEMSDB returns a slice of all rows with matching familyName and an error if any
func ReadOrdersByCustomerITEMSDB(familyName string) ([]Order, error) {

	var list []Order
	var s Order
	rows, err := db.Query("SELECT timeStamp, sellername, customername from items where customername=$1;", familyName)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		rows.Scan(&s.Date, &s.Shop, &s.FamilyName)
		list = append(list, s)
	}

	err = rows.Err()

	return list, err
}

//ReadOrdersToSellerITEMSDB returns all items orders sent to the seller
func ReadOrdersToSellerITEMSDB(sellerName string) ([]ShoppingList, error) {

	var list []ShoppingList
	var s ShoppingList
	rows, err := db.Query("SELECT timeStamp, customername, items from items where sellername=$1;", sellerName)
	if err != nil {
		return list, err
	}
	var items string
	for rows.Next() {
		rows.Scan(&s.Date, &s.FamilyName, &items)
		//split items and add to shopping list
		if items != "" {
			li := strings.Split(items, "\n")
			for _, v := range li {
				if v != "" {
					temp := strings.Split(v, "|")
					if len(temp) >= 2 {
						item := strings.Split(v, "|")[2]
						s.Items = append(s.Items, item)
					}
				}
			}
		}
		//add order to list
		list = append(list, s)
	}
	//rows does not return any error on scan, need to check explicitly
	err = rows.Err()

	return list, err
}

// ReadSellerDetailsITEMSDB returns the shop details
func ReadSellerDetailsITEMSDB(sellerName string) (Seller, error) {

	var shop Seller
	rows, err := getSellerRows(sellerName)
	if err != nil {
		return shop, err
	}
	for rows.Next() {
		rows.Scan(&shop.Name, &shop.Addr, &shop.OpenTime, &shop.CloseTime)
	}
	err = rows.Err()
	return shop, err

}
