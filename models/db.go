package models

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

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
	s = s + id + "|" + user + "|" + item + "\n"

	err = updateRow(s, familyName)
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
					v = ""
				}
				sNew += v + "\n"
			}
		}
	}

	//write row to db
	err = updateRow(sNew, familyName)

	return err
}

//ReadSellersDB returns a list of sellers
func ReadSellersDB() ([]Seller, error) {

	var listOfSellers []Seller
	var s Seller

	rows, err := getSellerRows()
	if err != nil {
		return listOfSellers, err
	}

	for rows.Next() {
		if err = rows.Scan(&s.ID, &s.Name, &s.Addr, &s.OpenTime, &s.CloseTime); err != nil {
			return listOfSellers, err
		}
		listOfSellers = append(listOfSellers, s)
	}
	return listOfSellers, err
}

//reads DB and sends items in row if found
// if row not found, creates one too
// if an error occurs, returns an empty string and the err
func getCustomerRow(familyName string) (string, error) {

	//query the family, create if does not exist
	row := db.QueryRow("SELECT items from customers where name = $1;", familyName)
	var err error
	var s string
	if err := row.Scan(&s); err != nil {

		if err == sql.ErrNoRows { // if row not found
			//create a new row
			err = nil
			_, err = db.Exec("INSERT into customers (name,items) VALUES ($1,$2);", familyName, "")
			//sanity check
			if err != nil {
				return "", err
			}
		}
	}
	return s, err
}

func getSellerRows() (*sql.Rows, error) {

	rows, err := db.Query("Select * from sellers")
	return rows, err
}
